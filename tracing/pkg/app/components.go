package app

import "go.opentelemetry.io/otel/trace"

type Components struct {
	RequestTracer    trace.Tracer
	RepositoryTracer trace.Tracer
}
