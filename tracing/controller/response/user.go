package response

import (
	"tracing/domain"
)

type Users []*User

type User struct {
	ID     int
	Mobile string
}

func MapperUser(domUser *domain.User) *User {
	if domUser == nil {
		return nil
	}
	respUser := new(User)
	respUser.ID = domUser.ID
	respUser.Mobile = domUser.Mobile
	return respUser
}

func MapperUsers(domainUsers domain.Users) Users {
	return mapper(domainUsers, MapperUser)
}
