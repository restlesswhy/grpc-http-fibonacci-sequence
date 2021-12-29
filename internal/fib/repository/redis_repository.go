package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
)

type redisRepo struct {
	redisClient *redis.Client
}

// Auth redis repository constructor
func NewRedisRepo(redisClient *redis.Client) fib.RedisRepository {
	return &redisRepo{redisClient: redisClient}
}

func (r *redisRepo) Add(x int32, y int64) error {
	return nil
}

func (r *redisRepo) CheckFib(x int32) (bool, error) {
	return false, nil
}