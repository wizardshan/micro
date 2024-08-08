package domain

import (
	"github.com/elliotchance/pie/v2"
	"time"
)

type Users []*User

func (dom Users) WithOnlineTimes(mapping map[int]UserOnlineTimes) Users {
	pie.Each(dom, func(item *User) {
		item.WithOnlineTimes(mapping[item.ID])
	})

	return dom
}

type User struct {
	ID       int
	Mobile   string
	Nickname string
	Bio      string

	CreateTime time.Time
	UpdateTime time.Time

	OnlineTimes    UserOnlineTimes
	OnlineDuration int
}

func (dom *User) WithOnlineTimes(onlineTimes UserOnlineTimes) *User {
	if onlineTimes == nil {
		return dom
	}
	dom.OnlineTimes = onlineTimes
	dom.OnlineDuration = onlineTimes.Duration()

	return dom
}
