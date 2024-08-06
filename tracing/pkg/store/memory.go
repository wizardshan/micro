package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/coocood/freecache"
)

type Memory struct {
	db *freecache.Cache
}

func NewMemory() Cache {
	c := new(Memory)
	cacheSize := 100 * 1024 * 1024
	c.db = freecache.NewCache(cacheSize)
	//debug.SetGCPercent(20)
	return c
}

func (c *Memory) Set(ctx context.Context, key string, value []byte, expireSeconds int) {
	err := c.db.Set([]byte(key), value, expireSeconds)
	if err != nil {
		panic(err)
	}
}

func (c *Memory) Get(ctx context.Context, key string) ([]byte, bool) {
	value, err := c.db.Get([]byte(key))
	if err == nil {
		return value, true
	}

	if errors.Is(err, freecache.ErrNotFound) {
		return nil, false
	}

	panic(err)
}

func (c *Memory) Lock(ctx context.Context, key string, expireSeconds int) bool {
	_, found, err := c.db.SetAndGet([]byte(fmt.Sprintf("lock:%s", key)), []byte(lockValue), expireSeconds)
	if err != nil {
		panic(err)
	}

	return !found
}
