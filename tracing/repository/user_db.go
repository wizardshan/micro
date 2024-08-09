package repository

import (
	"context"
	"tracing/domain"
	"tracing/repository/ent"
)

type UserDB struct {
	repo
}

func (repo *UserDB) FetchOne(ctx context.Context, optionFunc func(builder *ent.UserQuery)) *domain.User {
	return repo.fetchOne(ctx, repo.db, optionFunc).Mapper()
}

func (repo *UserDB) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserQuery)) *ent.User {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}

func (repo *UserDB) FetchMany(ctx context.Context, optionFunc func(builder *ent.UserQuery)) domain.Users {
	return repo.fetchMany(ctx, repo.db, optionFunc).Mapper()
}

func (repo *UserDB) fetchMany(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserQuery)) ent.Users {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.AllX(ctx)
}
