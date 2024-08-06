package repository

import (
	"context"
	"tracing/controller/request"
	"tracing/domain"
	"tracing/repository/ent"
	usercenter "tracing/repository/user-center"
)

type IUser interface {
	Fetch(ctx context.Context, id int) *domain.User
	FetchMany(ctx context.Context, query *request.Users) domain.Users
}

func NewUser(db *ent.Client, center *usercenter.Client) IUser {
	repo := new(UserProxy)
	repo.db = db
	repo.center = center
	return repo
}

type UserProxy struct {
	User
}
