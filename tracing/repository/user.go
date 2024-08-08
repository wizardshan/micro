package repository

import (
	"context"
	"time"
	"tracing/controller/request"
	"tracing/domain"
	"tracing/repository/ent"
	"tracing/repository/ent/user"
	"tracing/repository/service/bi"
)

type User struct {
	repo
	servBI *bi.Client
}

func (repo *User) Fetch(ctx context.Context, id int) *domain.User {
	domUser := repo.db.User.Query().WithComments().WithPosts().Where(user.ID(id)).FirstX(ctx).Mapper()

	onlineTimeQuery := &bi.OnlineTimeQuery{
		UserID:    1428754447,
		StartDate: time.Now().AddDate(0, 0, -1),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Terminal:  []int{1, 3, 4, 5, 6, 7},
	}

	onlineTimes := repo.servBI.OnlineTimes(ctx, onlineTimeQuery).Mapper()
	domUser.WithOnlineTimes(onlineTimes)

	return domUser
}

func (repo *User) FetchMany(ctx context.Context, query *request.Users) domain.Users {
	var entUser ent.Users = repo.db.User.Query().WithComments().WithPosts().AllX(ctx)
	domUsers := entUser.Mapper()

	userId := 1428754447
	onlineTimeQuery := &bi.OnlineTimeQuery{
		UserID:    userId,
		StartDate: time.Now().AddDate(0, 0, -1),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Terminal:  []int{1, 3, 4, 5, 6, 7},
	}

	onlineTimes := repo.servBI.OnlineTimes(ctx, onlineTimeQuery).Mapper()
	onlineTimesByUserIDMapping := map[int]domain.UserOnlineTimes{
		userId: onlineTimes,
	}
	domUsers.WithOnlineTimes(onlineTimesByUserIDMapping)

	return domUsers
}
