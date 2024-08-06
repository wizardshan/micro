package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	db *redis.Client
}

func NewRedis() Cache {
	c := new(Redis)
	c.db = redis.NewClient(&redis.Options{
		Addr:     "bjdxt.qcgameweb01.qiyi.redis:6522",
		Password: "LG12zOzf08xE", // no password set
		DB:       0,              // use default DB
	})

	return c
}

func (c *Redis) Set(ctx context.Context, key string, value []byte, expireSeconds int) {
	err := c.db.Set(ctx, key, value, time.Duration(expireSeconds)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func (c *Redis) Get(ctx context.Context, key string) ([]byte, bool) {

	value, err := c.db.Get(ctx, key).Bytes()
	if err == nil {
		return value, true
	}

	if errors.Is(err, redis.Nil) {
		return nil, false
	}

	panic(err)
}

func (c *Redis) Lock(ctx context.Context, key string, expireSeconds int) bool {
	success, err := c.db.SetNX(ctx, fmt.Sprintf("lock:%s", key), lockValue, time.Duration(expireSeconds)*time.Second).Result()
	if err != nil {
		panic(err)
	}

	return success
}
