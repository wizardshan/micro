package app

import (
	"go.opentelemetry.io/otel/trace"
	"tracing/pkg/store"
	"tracing/repository"
	"tracing/repository/ent"
)

type Tracer struct {
	Ctr trace.Tracer
	DB  trace.Tracer
}

type Components struct {
	DB     *ent.Client
	Cache  store.Cache
	Tracer *Tracer

	RepoUser repository.IUser
}
