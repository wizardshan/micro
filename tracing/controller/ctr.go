package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tracing/pkg/store"
)

type Response struct {
	Code    int
	Message string
	Success bool
	Router  string
	Data    any
}

func (resp *Response) Marshal() []byte {
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return b
}

type ctr struct {
	cache *store.Cache
}

func cache[T any](c *gin.Context, cache *store.Cache, key string, fetch func() (T, error)) (resp T, err error) {

	if data, found := cache.Get(c.Request.Context(), key); found {
		c.Header("Cache-Key", key)
		c.Header("Cached", "true")
		var resp T
		err = json.Unmarshal(data, &resp)
		return resp, err
	}

	c.Header("Cached", "false")
	if resp, err = fetch(); err != nil {
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return
	}
	cache.Set(c.Request.Context(), key, data, 10)
	return resp, nil
}
