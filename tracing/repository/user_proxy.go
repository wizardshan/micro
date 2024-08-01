package repository

import (
	"context"
	"tracing/pkg/app"
	"tracing/repository/ent"
)

type IUser interface {
	Fetch(ctx context.Context, id int) *ent.User
}

func NewUser(db *ent.Client, components *app.Components) IUser {
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