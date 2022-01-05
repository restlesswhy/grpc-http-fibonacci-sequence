//go:generate mockgen -source redis_repository.go -destination mock/redis.go -package mock
package fib

import "context"

type RedisRepository interface {
	CheckFib(ctx context.Context, key string) (string, bool, error)	
	Add(ctx context.Context, key string, value string) error
}

