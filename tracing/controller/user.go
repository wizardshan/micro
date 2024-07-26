package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
	"tracing/pkg/app"
	"tracing/repository"
)

type User struct {
	repo   *repository.User
	tracer trace.Tracer
}

func NewUser(repo *repository.User, components *app.Components) *User {
	ctr := new(User)
	ctr.repo = repo
	ctr.tracer = components.RepositoryTracer
	return ctr
}

func (ctr *User) One(c *gin.Context) {
	id := 1

	_, span := ctr.tracer.Start(c.Request.Context(), "FetchName", trace.WithAttributes(attribute.Int("user.id", id)))
	defer span.End()

	fmt.Printf("%s\n", span.SpanContext().TraceID().String())
	// add start event
	span.AddEvent("start to get user",
		trace.WithTimestamp(time.Now()),
	)

	name, err := ctr.repo.FetchName(c.Request.Context(), id)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		c.String(http.StatusOK, err.Error())
		return
	}

	// set user info in span's attributes
	span.SetAttributes(attribute.String("user.name", name))
	// add end event
	span.AddEvent("end to get user",
		trace.WithTimestamp(time.Now()),
		trace.WithAttributes(attribute.String("user.name", name)),
	)
	span.SetStatus(codes.Ok, "")

	c.String(http.StatusOK, name)
}
