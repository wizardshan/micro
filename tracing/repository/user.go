package repository

import (
	"context"
	"time"
	"tracing/repository/ent"
	"tracing/repository/ent/user"
)

type User struct {
	repo
}

func (repo *User) Fetch(ctx context.Context, id int) *ent.User {
	time.Sleep(500 * time.Millisecond)

	return repo.db.User.Query().WithComments().WithPosts().Where(user.ID(id)).FirstX(ctx)
}
