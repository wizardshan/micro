package response

import (
	"tracing/domain"
)

type Users []*User

type User struct {
	ID     int
	Mobile string

	OnlineTimes    UserOnlineTimes
	OnlineDuration int
}

func MapperUser(domUser *domain.User) *User {
	if domUser == nil {
		return nil
	}
	respUser := new(User)
	respUser.ID = domUser.ID
	respUser.Mobile = domUser.Mobile

	respUser.OnlineDuration = domUser.OnlineDuration
	respUser.OnlineTimes = MapperUserOnlineTimes(domUser.OnlineTimes)

	return respUser
}

func MapperUsers(domainUsers domain.Users) Users {
	return mapper(domainUsers, MapperUser)
}
