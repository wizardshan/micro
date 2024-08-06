package repository

import (
	"context"
	"tracing/controller/request"
	"tracing/domain"
	"tracing/repository/ent"
	"tracing/repository/ent/user"
	usercenter "tracing/repository/user-center"
)

type User struct {
	repo
	center *usercenter.Client
}

func (repo *User) Fetch(ctx context.Context, id int) *domain.User {
	return repo.db.User.Query().WithComments().WithPosts().Where(user.ID(id)).FirstX(ctx).Mapper()
}

func (repo *User) FetchMany(ctx context.Context, query *request.Users) domain.Users {
	var entUser ent.Users = repo.db.User.Query().WithComments().WithPosts().Where(user.ID(query.ID)).AllX(ctx)

	repo.center.FetchMany(ctx, []int{1428754447})

	return entUser.Mapper()
}
