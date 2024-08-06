package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/coocood/freecache"
)

const (
	levelMemory = iota
	levelRedis
	levelMemoryRedis
)

type Cache struct {
	memory *freecache.Cache
	level  int
}

func New() *Cache {
	c := new(Cache)
	cacheSize := 100 * 1024 * 1024
	c.memory = freecache.NewCache(cacheSize)
	//debug.SetGCPercent(20)
	return c
}

func (c *Cache) Set(ctx context.Context, key string, value []byte, expireSeconds int) {
	err := c.memory.Set([]byte(key), value, expireSeconds)
	if err != nil {
		panic(err)
	}
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, bool) {
	value, err := c.memory.Get([]byte(key))
	if err == nil {
		return value, true
	}

	if errors.Is(err, freecache.ErrNotFound) {
		return nil, false
	}

	panic(err)
}

func (c *Cache) Lock(ctx context.Context, key string, expireSeconds int) bool {
	_, found, err := c.memory.SetAndGet([]byte(fmt.Sprintf("lock:%s", key)), []byte("locking"), expireSeconds)
	if err != nil {
		panic(err)
	}

	return !found
}
