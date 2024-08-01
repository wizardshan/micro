package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"tracing/repository"
)

type User struct {
	repo repository.IUser
}

func (ctr *User) One(c *gin.Context) (any, error) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	time.Sleep(500 * time.Millisecond)
	return ctr.repo.Fetch(c.Request.Context(), id), nil
}
