package controller

import (
	"github.com/gin-gonic/gin"
	"tracing/pkg/app"
	"tracing/repository"
)

type IUser interface {
	One(c *gin.Context) (any, error)
}

func NewUser(repo repository.IUser, components *app.Components) IUser {
	ctr := new(UserProxy)
	ctr.repo = repo
	return ctr
}

type UserProxy struct {
	User
}
