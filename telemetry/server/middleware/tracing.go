package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/url"
	"strings"
	"time"
)

func Tracing(ctx context.Context, appName string, endPoint string) gin.HandlerFunc {

	provider, err := initProvider(ctx, appName, endPoint)
	if err != nil {
		panic(err)
	}

	opts := []otelgin.Option{
		otelgin.WithTracerProvider(provider),
		otelgin.WithPropagators(propagation.TraceContext{}), // 从上游获取traceparent, tracestate
	}

	return otelgin.Middleware(appName, opts...)
}

func initProvider(ctx context.Context, appName string, endPoint string) (trace.TracerProvider, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(appName),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	exporter, err := traceExporter(ctx, endPoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	return tracerProvider, nil
}

func traceExporter(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	ur, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	var exporter *otlptrace.Exporter

	switch strings.ToLower(ur.Scheme) {
	case "http", "https":
		exporter, err = httpExporter(ctx, ur.Host)
		if err != nil {
			return nil, err
		}

	case "grpc":
		fallthrough
	default:
		exporter, err = grpcExpoter(ctx, ur.Host)
		if err != nil {
			return nil, err
		}
	}

	return exporter, nil
}

func grpcExpoter(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	// addr := strings.TrimLeft(endpoint, "grpc://")

	conn, err := grpc.DialContext(ctx, endpoint,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		// grpc.WithTimeout(5*time.Second),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithGRPCConn(conn),
		// otlptracegrpc.WithHeaders(
		// 	map[string]string{
		// 		"authorization": BearerAuthToken,
		// 		"Authorization": BearerAuthToken,
		// 	},
		// ),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	return traceExporter, nil
}

func httpExporter(ctx context.Context, endPoint string) (*otlptrace.Exporter, error) {

	// endpoint = strings.TrimPrefix(endpoint, "https://")
	// endpoint = strings.TrimPrefix(endpoint, "http://")

	opts := []otlptracehttp.Option{
		otlptracehttp.WithTimeout(5 * time.Second),
		otlptracehttp.WithEndpoint(endPoint),
		otlptracehttp.WithInsecure(),
		// otlptracehttp.WithHeaders(
		// 	map[string]string{
		// 		"authorization": BearerAuthToken,
		// 		"Authorization": BearerAuthToken,
		// 	},
		// ),
	}

	trace, err := otlptracehttp.New(ctx, opts...)

	return trace, err
}
