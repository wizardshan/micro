package request

import "fmt"

const cacheKeyPrefix = "Ctr"

type User struct {
	Mobile string
}

func (req *User) CacheKey(id int) string {
	return fmt.Sprintf("%s:User:ID:%d", cacheKeyPrefix, id)
}

type Users struct {
	ID     int
	Mobile string
}

func (req *Users) CacheKey() string {
	return fmt.Sprintf("%s:Users:ID:%d:Mobile:%s", cacheKeyPrefix, req.ID, req.Mobile)
}
