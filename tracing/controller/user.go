package controller

import (
	"github.com/gin-gonic/gin"
	"tracing/controller/request"
	"tracing/controller/response"
	"tracing/repository"
)

type User struct {
	repo repository.IUser
}

func (ctr *User) One(c *gin.Context, query *request.User, id int) *response.User {
	domUser := ctr.repo.Fetch(c.Request.Context(), id)
	return response.MapperUser(domUser)
}

func (ctr *User) Many(c *gin.Context, query *request.Users) response.Users {
	domUsers := ctr.repo.FetchMany(c.Request.Context(), query)
	return response.MapperUsers(domUsers)
}
