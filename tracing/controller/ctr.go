package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int
	Message string
	Success bool
	Router  string
	Data    any
}

type HandlerFunc func(c *gin.Context) (any, error)

type Handler struct {
}

func (handler *Handler) Wrapper(handlerFunc HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		router := c.FullPath()

		data, err := handlerFunc(c)
		if err == nil {
			c.JSON(http.StatusOK, Response{Code: 0, Message: "", Success: true, Router: router, Data: data})
			return
		}

		c.JSON(http.StatusOK, Response{Code: 1, Message: err.Error(), Success: false, Router: router, Data: nil})
	}
}
