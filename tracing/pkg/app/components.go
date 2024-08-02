package app

import "go.opentelemetry.io/otel/trace"

type Tracer struct {
	Ctr trace.Tracer
	DB  trace.Tracer
}

type Components struct {
	Tracer *Tracer
}
