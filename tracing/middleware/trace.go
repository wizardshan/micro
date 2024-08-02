package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type Filter func(*http.Request) bool

func Trace(tracer trace.Tracer) gin.HandlerFunc {

	var filters []Filter

	return func(c *gin.Context) {
		for _, f := range filters {
			if !f(c.Request) {
				// Serve the request to the next middleware if a filter rejects the request.
				c.Next()
				return
			}
		}
		savedCtx := c.Request.Context()
		defer func() {
			c.Request = c.Request.WithContext(savedCtx)
		}()

		header := c.Request.Header
		header.Set("traceParent", c.DefaultQuery("traceParent", ""))

		propagators := otel.GetTextMapPropagator()
		ctx := propagators.Extract(savedCtx, propagation.HeaderCarrier(header))

		opts := []trace.SpanStartOption{
			trace.WithAttributes(attribute.String("http.client_ip", c.ClientIP())),
			trace.WithAttributes(attribute.String("http.target", c.Request.URL.Path)),
			trace.WithAttributes(attribute.String("user_agent", c.Request.UserAgent())),
			trace.WithSpanKind(trace.SpanKindServer),
		}
		spanName := fmt.Sprintf("Router %s %s", c.Request.Method, c.FullPath())
		ctx, span := tracer.Start(ctx, spanName, opts...)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		status := c.Writer.Status()
		span.SetStatus(ServerStatus(status))
		span.SetAttributes(attribute.Int("http.status_code", status))

		if len(c.Errors) > 0 {
			span.SetAttributes(attribute.String("gin.errors", c.Errors.String()))
		}
	}
}

func ServerStatus(code int) (codes.Code, string) {
	if code < 100 || code >= 600 {
		return codes.Error, fmt.Sprintf("Invalid HTTP status code %d", code)
	}
	if code >= 500 {
		return codes.Error, ""
	}
	return codes.Unset, ""
}
