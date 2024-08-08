package app

import (
	"go.opentelemetry.io/otel/trace"
	"tracing/pkg/store"
	"tracing/repository/ent"
)

type Tracer struct {
	Ctr trace.Tracer
	DB  trace.Tracer
}

type ServerInfo struct {
	Host    string
	Source  int
	SignKey string
}

type Servers struct {
	BI *ServerInfo
}

type Components struct {
	DB     *ent.Client
	Cache  store.Cache
	Tracer *Tracer

	Servers *Servers
}
