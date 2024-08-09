package repository

import (
	"context"
	"tracing/domain"
	"tracing/pkg/timepoint"
	"tracing/repository/ent"
	"tracing/repository/service/game"
)

func NewUser(db *ent.Client, servBI *game.BI, servPayment *game.Payment) *User {
	repo := new(User)
	repo.db = db
	repo.servBI = servBI
	repo.servPayment = servPayment
	return repo
}

type User struct {
	UserDB
	servBI      *game.BI
	servPayment *game.Payment
}

func (repo *User) FetchOne(ctx context.Context, queryFunc func(builder *ent.UserQuery), options ...UserQueryOption) *domain.User {
	domUser := repo.UserDB.FetchOne(ctx, queryFunc)

	if len(options) == 0 {
		return domUser
	}

	uq := &UserQuery{
		withOnlineTime: false,
		withOrder:      false,
	}
	for _, opt := range options {
		opt(uq)
	}

	if uq.withOnlineTime {
		domUser.WithOnlineTimes(repo.OnlineTimes(ctx, domUser.ID))
	}

	if uq.withOrder {
		repo.Orders(ctx, domUser.ID)
	}

	return domUser
}

func (repo *User) Orders(ctx context.Context, ID int) domain.UserOnlineTimes {
	return repo.servPayment.Orders(ctx, func(q *game.OrderQuery) {
		q.UserID = ID
		q.StartTime = timepoint.StartOfMonth()
		q.EndTime = timepoint.EndOfMonth()
	}).Mapper()
}

func (repo *User) OnlineTimes(ctx context.Context, ID int) domain.UserOnlineTimes {
	return repo.servBI.OnlineTimes(ctx, func(q *game.OnlineTimeQuery) {
		q.UserID = ID
		q.StartDate = timepoint.StartOfDay()
		q.EndDate = timepoint.EndOfDay()
	}).Mapper()
}

func (repo *User) FetchMany(ctx context.Context, queryFunc func(builder *ent.UserQuery), options ...UserQueryOption) domain.Users {
	domUsers := repo.UserDB.FetchMany(ctx, queryFunc)

	userId := 1428754447
	onlineTimes := repo.servBI.OnlineTimes(ctx, func(q *game.OnlineTimeQuery) {
		q.UserID = userId
		q.StartDate = timepoint.StartOfDay()
		q.EndDate = timepoint.EndOfDay()
	}).Mapper()

	onlineTimesByUserIDMapping := map[int]domain.UserOnlineTimes{
		userId: onlineTimes,
	}
	domUsers.WithOnlineTimes(onlineTimesByUserIDMapping)

	return domUsers
}
