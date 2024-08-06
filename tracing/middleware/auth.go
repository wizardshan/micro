package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	c.UserID = 111
	c.Next()
}
