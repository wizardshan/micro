package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tracing/controller/request"
	"tracing/controller/response"
	"tracing/pkg/app"
	"tracing/repository"
)

type UserProxy struct {
	ctr
	User
}

func NewUser(user repository.IUser, components *app.Components) *UserProxy {
	ctr := new(UserProxy)
	ctr.repo = user
	ctr.cache = components.Cache
	return ctr
}

func (ctr *UserProxy) One(c *gin.Context) (any, error) {
	request := new(request.User)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	return cache(c, ctr.cache, request.CacheKey(id), func() (*response.User, error) {
		return ctr.User.One(c, request, id), nil
	})
}

func (ctr *UserProxy) Many(c *gin.Context) (any, error) {
	request := new(request.Users)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	return cache(c, ctr.cache, request.CacheKey(), func() (response.Users, error) {
		return ctr.User.Many(c, request), nil
	})
}
