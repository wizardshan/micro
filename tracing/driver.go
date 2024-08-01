package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Driver struct {
	dialect.Driver
	tracer trace.Tracer
}

func NewDriver(drv dialect.Driver, tracer trace.Tracer) *Driver {
	return &Driver{
		Driver: drv,
		tracer: tracer,
	}
}

func (d *Driver) Query(ctx context.Context, query string, argsRaw, v any) error {

	args, _ := argsRaw.([]any)

	var attrs []string
	for _, v := range args {
		attrs = append(attrs, fmt.Sprintf("%v", v))
	}

	ctx, span := d.tracer.Start(ctx, query, trace.WithAttributes(attribute.StringSlice("query.args", attrs)))
	defer span.End()

	err := d.Driver.Query(ctx, query, argsRaw, v)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	span.SetStatus(codes.Ok, "")

	return nil
}
