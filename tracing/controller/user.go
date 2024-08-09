package controller

import (
	"github.com/gin-gonic/gin"
	"tracing/controller/request"
	"tracing/controller/response"
	"tracing/repository"
	"tracing/repository/ent"
	"tracing/repository/ent/user"
)

type User struct {
	repo *repository.User
}

func (ctr *User) One(c *gin.Context, query *request.User, id int) *response.User {

	domUser := ctr.repo.FetchOne(c.Request.Context(), func(query *ent.UserQuery) {
		query.Where(user.ID(id))
	}, repository.UserQueryWithOrder(true))

	return response.MapperUser(domUser)
}

func (ctr *User) Many(c *gin.Context, query *request.Users) response.Users {
	domUsers := ctr.repo.FetchMany(c.Request.Context(), func(query *ent.UserQuery) {
	})
	return response.MapperUsers(domUsers)
}
