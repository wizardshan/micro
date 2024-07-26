package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"tracing/controller"
	"tracing/middleware"
	"tracing/pkg/app"
	"tracing/repository"
)

func main() {

	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		panic(err.Error())
	}

	requestTracer := tp.Tracer(
		"request",
		trace.WithInstrumentationVersion("1.0"),
	)

	repositoryTracer := tp.Tracer(
		"repository",
		trace.WithInstrumentationVersion("1.0"),
	)

	components := &app.Components{
		RequestTracer:    requestTracer,
		RepositoryTracer: repositoryTracer,
	}

	engine := gin.New()
	engine.Use(middleware.Middleware(requestTracer))

	repoUser := repository.NewUser()
	ctrUser := controller.NewUser(repoUser, components)
	engine.GET("/user/name", ctrUser.One)

	engine.Run()
}
