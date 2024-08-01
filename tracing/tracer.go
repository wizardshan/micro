package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.12.0"
)

var ServiceName = "UserServer"

func initTracer(ctx context.Context) (*trace.TracerProvider, error) {
	// http导出器
	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(ServiceName),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String("prod"),
		),
	)
	if err != nil {
		return nil, err
	}

	// SimpleSpanProcessor 立即将结束的 Span 转发给导出器
	// BatchSpanProcessor 批量处理并批量发送
	bsp := trace.NewBatchSpanProcessor(exp)

	// 输出到控制台
	//traceExporter, err := stdouttrace.New(
	//	stdouttrace.WithPrettyPrint())
	//if err != nil {
	//	return nil, err
	//}

	/*
		AlwaysSample(): 全部采样
		NeverSample(): 全部丢弃
		TraceIDRatioBased(fraction float64): 设置采样率
		ParentBased(root Sampler, samplers ...ParentBasedSamplerOption): 基于parent span 设置采样策略
	*/
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		//sdktrace.WithSampler(trace.TraceIDRatioBased(0.5)), // 采样率为50%
		//sdktrace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.5))), // 根据parent span采样策略
		trace.WithSpanProcessor(bsp),
		//trace.WithSpanProcessor(trace.NewSimpleSpanProcessor(traceExporter)),
		trace.WithResource(res),
	)

	// 设置全局TraceProvider
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{}) // 使用TraceContext在下游Inject和上游Extract来打通服务间调用链路
	return tp, nil
}
