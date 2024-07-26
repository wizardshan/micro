package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/middleware"
)

var appName = "ServerApp"
var endPoint = "http://localhost:4318"

func main() {
	ctx := context.Background()

	r := gin.Default()
	r.Use(middleware.Tracing(ctx, appName, endPoint))
	r.Use(middleware.Tracing(ctx, appName, endPoint))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Traceparent": c.GetHeader("Traceparent"),
		})
	})

	r.Run(":3001")
}
