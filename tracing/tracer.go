package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var ServiceName = "order"

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// http导出器
	exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(ServiceName),
		),
	)
	// SimpleSpanProcessor 立即将结束的 Span 转发给导出器
	// BatchSpanProcessor 批量处理并批量发送
	bsp := sdktrace.NewBatchSpanProcessor(exp)

	/*
		AlwaysSample(): 全部采样
		NeverSample(): 全部丢弃
		TraceIDRatioBased(fraction float64): 设置采样率
		ParentBased(root Sampler, samplers ...ParentBasedSamplerOption): 基于parent span 设置采样策略
	*/
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		//sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.5)), // 采样率为50%
		//sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.5))), // 根据parent span采样策略
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
	)

	// 设置全局TraceProvider
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{}) // 使用TraceContext在下游Inject和上游Extract来打通服务间调用链路
	return tp, nil
}
