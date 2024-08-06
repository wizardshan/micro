package domain

import "time"

type Users []*User

type User struct {
	ID       int
	Mobile   string
	Nickname string
	Bio      string

	CreateTime time.Time
	UpdateTime time.Time
}
