package fib

import "context"

type RedisRepository interface {
	CheckFib(ctx context.Context, key string) (int64, bool, error)	
	Add(ctx context.Context, key string, value int64) error
}

