package repository

import (
	"go.opentelemetry.io/otel/trace"
	"tracing/repository/ent"
)

type repo struct {
	db     *ent.Client
	tracer trace.Tracer
}
