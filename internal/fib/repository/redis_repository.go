package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	// "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

type redisRepo struct {
	cfg *config.Config
	redisClient *redis.Client
}

// Auth redis repository constructor
func NewRedisRepo(redisClient *redis.Client, cfg *config.Config) fib.RedisRepository {
	return &redisRepo{redisClient: redisClient, cfg: cfg}
}

func (r *redisRepo) Add(ctx context.Context, key string, value int64) error {
	if err := r.redisClient.Set(ctx, key, value, r.cfg.Redis.FibTTL).Err(); err != nil {
		return errors.Wrap(err, "redisRepo.Add.redisClient.Set")
	}

	return nil
}

func (r *redisRepo) CheckFib(ctx context.Context, key string) (int64, bool, error) {
	var isExist bool
	var res int64
	var err error
	
	if res, err = r.redisClient.Get(ctx, key).Int64(); err == redis.Nil {
		return 0, isExist, nil
	} else if err != nil {
		return 0, isExist, errors.Wrap(err, "redisRepo.Add.redisClient.Set")
	}

	isExist = true

	return res, isExist, nil
}