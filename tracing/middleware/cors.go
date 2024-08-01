package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(gc *gin.Context) {
		gc.Header("Access-Control-Allow-Origin", "*")
		gc.Header("Access-Control-Allow-Headers", "Content-Type, Traceparent")
		gc.Header("Access-Control-Allow-Methods", "POST, GET,PUT,DELETE, OPTIONS")
		gc.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		gc.Header("Access-Control-Allow-Credentials", "true")

		method := gc.Request.Method
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			gc.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		gc.Next()
	}
}
