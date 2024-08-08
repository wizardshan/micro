package repository

import (
	"context"
	"tracing/controller/request"
	"tracing/domain"
	"tracing/repository/ent"
	"tracing/repository/service/bi"
)

type IUser interface {
	Fetch(ctx context.Context, id int) *domain.User
	FetchMany(ctx context.Context, query *request.Users) domain.Users
}

func NewUser(db *ent.Client, servBI *bi.Client) IUser {
	repo := new(UserProxy)
	repo.db = db
	repo.servBI = servBI
	return repo
}

type UserProxy struct {
	User
}
