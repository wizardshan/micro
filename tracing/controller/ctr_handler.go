package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"tracing/pkg/cache"
)

type Filter func(*http.Request) bool

var filters []Filter

type HandlerFunc func(c *gin.Context) (any, error)

type Handler struct {
	tracer trace.Tracer
	cache  *cache.Cache
}

func NewHandler(tracer trace.Tracer, cache *cache.Cache) *Handler {
	handler := new(Handler)
	handler.tracer = tracer
	handler.cache = cache
	return handler
}

func (h *Handler) Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("cacheKey", []byte("ctr-"+c.FullPath()))

		c.JSON(http.StatusOK, h.traceWrapper(c, handlerFunc))
	}
}

func (h *Handler) handle(c *gin.Context, handlerFunc HandlerFunc) *Response {
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
			return h.cacheWrapper(c, handlerFunc)
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
	resp := h.cacheWrapper(c, handlerFunc)

	span.SetStatus(codes.Ok, resp.Message)
	if resp.Code != 0 {
		span.SetStatus(codes.Error, resp.Message)
		span.SetAttributes(attribute.String("resp", string(resp.bytes())))
	}

	return resp
}

func (h *Handler) cacheWrapper(c *gin.Context, handlerFunc HandlerFunc) *Response {
	cacheKey, exist := c.Get("cacheKey")
	if !exist {
		return h.handle(c, handlerFunc)
	}

	key, ok := cacheKey.([]byte)
	if ok {
		if data, found := h.cache.Get(c.Request.Context(), key); found {
			var resp Response
			if err := json.Unmarshal(data, &resp); err == nil {
				resp.Cache = true
				return &resp
			}

		}
	}

	resp := h.handle(c, handlerFunc)
	h.cache.Set(c.Request.Context(), key, resp.bytes(), 10)
	return resp
}
