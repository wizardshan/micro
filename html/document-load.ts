import {
    BatchSpanProcessor,
    WebTracerProvider,
} from '@opentelemetry/sdk-trace-web';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http';
import {ZoneContextManager} from "@opentelemetry/context-zone";
import { W3CTraceContextPropagator } from '@opentelemetry/core';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
import { XMLHttpRequestInstrumentation } from '@opentelemetry/instrumentation-xml-http-request';

const exporter = new OTLPTraceExporter({
    url: 'http://test.telemetry.woa.com:55681/v1/traces',
})

const provider = new WebTracerProvider();
provider.addSpanProcessor(new BatchSpanProcessor(exporter));
provider.resource.attributes['tps.tenant.id'] = 'test';
provider.resource.attributes['service.name'] = "test-service";

provider.register({
    contextManager: new ZoneContextManager(),
    propagator: new W3CTraceContextPropagator(),
});


registerInstrumentations({
    instrumentations: [
        new XMLHttpRequestInstrumentation({
            ignoreUrls: [/localhost:8090\/sockjs-node/],
            propagateTraceHeaderCorsUrls: [
                `/http:\/\/localhost:7777.*/`
            ],
        }),
    ],
});