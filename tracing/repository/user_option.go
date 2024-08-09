package repository

type UserQuery struct {
	withOnlineTime bool
	withOrder      bool
}

type UserQueryOption func(*UserQuery)

func UserQueryWithOnlineTime(v bool) UserQueryOption {
	return func(uq *UserQuery) {
		uq.withOnlineTime = v
	}
}

func UserQueryWithOrder(v bool) UserQueryOption {
	return func(uq *UserQuery) {
		uq.withOrder = v
	}
}
