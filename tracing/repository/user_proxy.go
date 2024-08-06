package repository

import (
	"context"
	"tracing/controller/request"
	"tracing/domain"
	"tracing/repository/ent"
)

type IUser interface {
	Fetch(ctx context.Context, id int) *domain.User
	FetchMany(ctx context.Context, query *request.Users) domain.Users
}

func NewUser(db *ent.Client) IUser {
	repo := new(UserProxy)
	repo.db = db
	return repo
}

type UserProxy struct {
	User
}

//func (repo *UserProxy) Fetch(ctx context.Context, id int) (entityUser *entity.User) {
//
//	repo.trace(ctx, func() {
//		entityUser = repo.User.Fetch(ctx, id)
//	})
//	return
//}

//func (repo *UserProxy) trace(ctx context.Context, processor func()) {
//	ctx, span := repo.tracer.Start(ctx, "UserRepo Fetch", trace.WithAttributes(attribute.Int("user.id", id)))
//	defer span.End()
//
//	processor()
//
//	//entityUser = repo.User.Fetch(ctx, id)
//	//if err != nil {
//	//	span.SetStatus(codes.Error, err.Error())
//	//	return
//	//}
//
//	//span.SetAttributes(attribute.String("user.nickname", entityUser.Nickname))
//	span.SetStatus(codes.Ok, "")
//	return
//}
