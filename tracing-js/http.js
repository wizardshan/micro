const { Resource } = require('@opentelemetry/resources');
const {SimpleSpanProcessor, WebTracerProvider} = require('@opentelemetry/sdk-trace-web');
const {ZoneContextManager} = require('@opentelemetry/context-zone');
const {OTLPTraceExporter} = require('@opentelemetry/exporter-trace-otlp-http');
const {context, propagation, trace, SpanKind} = require('@opentelemetry/api');
//const { W3CTraceContextPropagator } = require('@opentelemetry/core');
const {ConsoleSpanExporter} = require("@opentelemetry/sdk-trace-base");
const {
    SEMRESATTRS_SERVICE_NAME,
    SEMRESATTRS_SERVICE_VERSION,
    SEMRESATTRS_DEPLOYMENT_ENVIRONMENT,
} = require('@opentelemetry/semantic-conventions');

const traceServerUrl = 'http://test.com/v1/traces'
const serverHost = 'http://127.0.0.1:8080'

const exporter = new OTLPTraceExporter({
    url: traceServerUrl,
})

const resource = Resource.default().merge(
    new Resource({
        [SEMRESATTRS_SERVICE_NAME]: 'UserWeb',
        [SEMRESATTRS_SERVICE_VERSION]: '1.0.0',
        [SEMRESATTRS_DEPLOYMENT_ENVIRONMENT]: 'prod'
    }),
);

const provider = new WebTracerProvider({
    resource: resource,
});

// Note: For production consider using the "BatchSpanProcessor" to reduce the number of requests
// to your exporter. Using the SimpleSpanProcessor here as it sends the spans immediately to the
// exporter without delay
provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter())); //控制台上自动打印span
provider.addSpanProcessor(new SimpleSpanProcessor(exporter)); // 立刻发送,没有批处理操作，可用于调试学习
//provider.addSpanProcessor(new BatchSpanProcessor(exporter)); //批量处理跨度并批量发送它们

provider.register({
    contextManager: new ZoneContextManager(),
    //propagator: new W3CTraceContextPropagator(),
});

const tracer = provider.getTracer('JSClient');

const prepareClickEvent = () => {
    const id = 1
    const url = serverHost + '/user/' + id;

    const clickHandler = (fetchFn) => () => {
        const span = tracer.startSpan('', { kind: SpanKind.CLIENT});

        context.with(trace.setSpan(context.active(), span), () => {

            // 手动生成span，需要手动在header增加Traceparent参数
            const output = {};
            propagation.inject(context.active(), output);
            const { traceparent } = output; // version - traceid - parentid/spanid - traceflags
            console.log(traceparent)
            span.setAttribute("user.id", id)

            span.addEvent("requestStart")
            const method = 'GET'
            fetchFn(url, method, traceparent).then((response) => {
                span.updateName(method + " " +response.Router)
                span.addEvent("responseEnd")
                //trace.getSpan(context.active()).addEvent('fetching-single-span-completed');
                // 错误异常 span.setStatus()
                span.end();
            });
        });
    };

    const element1 = document.getElementById('button1');
    const element2 = document.getElementById('button2');
    element1.addEventListener('click', clickHandler(fetchData));
    element2.addEventListener('click', clickHandler(xhrData));
}

const fetchData = (url, method, traceparent) => fetch(url, {
    method: method,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        'Traceparent': traceparent,
    },
}).then(res=>res.json())

const xhrData = (url, method, traceparent) => new Promise((resolve, reject) => {
    const req = new XMLHttpRequest();
    req.open(method, url, true);
    req.setRequestHeader('Content-Type', 'application/json');
    req.setRequestHeader('Accept', 'application/json');
    req.setRequestHeader('Traceparent', traceparent);
    req.onload = () => {
        resolve();
    };
    req.onerror = () => {
        reject();
    };
    req.send();
});

window.addEventListener('load', prepareClickEvent);
