package repository

import (
	"context"
	"time"
)

type User struct {
}

func NewUser() *User {
	repo := new(User)
	return repo
}

func (repo *User) FetchName(ctx context.Context, id int) (string, error) {
	time.Sleep(1 * time.Second)
	return "admin", nil
}
