package cache

import (
	"context"
	"errors"
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

func (c *Cache) Set(ctx context.Context, key, value []byte, expireSeconds int) {
	err := c.memory.Set(key, value, expireSeconds)
	if err != nil {
		panic(err)
	}
}

func (c *Cache) Get(ctx context.Context, key []byte) ([]byte, bool) {
	value, err := c.memory.Get(key)
	if err == nil {
		return value, true
	}

	if errors.Is(err, freecache.ErrNotFound) {
		return nil, false
	}

	panic(err)
}
