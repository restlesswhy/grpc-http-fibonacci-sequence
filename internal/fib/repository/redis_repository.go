package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
	// "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
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

func (r *redisRepo) Add(ctx context.Context, key string, value string) error {
	if err := r.redisClient.Set(ctx, key, value, r.cfg.Redis.FibTTL).Err(); err != nil {
		return errors.Wrap(err, "redisRepo.Add.redisClient.Set")
	}
	// logger.Infof("added: %v, value: %v", key, value)
	return nil
}

func (r *redisRepo) CheckFib(ctx context.Context, key string) (string, bool, error) {
	var isExist bool
	var res string
	var err error
	
	if res, err = r.redisClient.Get(ctx, key).Result(); err == redis.Nil {
		logger.Info("exist: ", isExist)
		return "", isExist, nil
	} else if err != nil {
		return "", isExist, errors.Wrap(err, "redisRepo.Add.redisClient.Get")
	}
		
	isExist = true
	return res, isExist, nil
}