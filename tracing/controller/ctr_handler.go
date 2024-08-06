package controller

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

var filters []Filter

type HandlerFunc func(c *gin.Context) (any, error)

type Handler struct {
	tracer trace.Tracer
}

func NewHandler(tracer trace.Tracer) *Handler {
	handler := new(Handler)
	handler.tracer = tracer
	return handler
}

func (h *Handler) Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, h.traceWrapper(c, handlerFunc))
	}
}

func (h *Handler) handle(c *gin.Context, handlerFunc HandlerFunc) *Response {

	// 在这里处理panic

	data, err := handlerFunc(c)
	router := c.FullPath()
	if err == nil {
		return &Response{Code: 0, Message: "", Success: true, Router: router, Data: data}
	}

	return &Response{Code: 1, Message: err.Error(), Success: false, Router: router, Data: nil}
}

func (h *Handler) traceWrapper(c *gin.Context, handlerFunc HandlerFunc) *Response {
	for _, f := range filters {
		if !f(c.Request) {
			return h.handle(c, handlerFunc)
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
	spanName := fmt.Sprintf("Ctr %s %s", c.Request.Method, c.FullPath())
	ctx, span := h.tracer.Start(ctx, spanName, opts...)
	defer span.End()

	c.Request = c.Request.WithContext(ctx)
	resp := h.handle(c, handlerFunc)

	span.SetStatus(codes.Ok, resp.Message)
	cached := c.Writer.Header().Get("Cached")
	span.SetAttributes(attribute.String("Cached", cached))
	if c.Writer.Header().Get("Cached") == "true" {
		span.SetAttributes(attribute.String("Cache-Key", c.Writer.Header().Get("Cache-Key")))
	}
	if resp.Code != 0 {
		span.SetStatus(codes.Error, resp.Message)
		span.SetAttributes(attribute.String("resp", string(resp.Marshal())))
	}

	return resp
}
