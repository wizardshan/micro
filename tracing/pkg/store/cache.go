package store

import "context"

const lockValue = "locking"

type Cache interface {
	Set(ctx context.Context, key string, value []byte, expireSeconds int)
	Get(ctx context.Context, key string) ([]byte, bool)
	Lock(ctx context.Context, key string, expireSeconds int) bool
}
