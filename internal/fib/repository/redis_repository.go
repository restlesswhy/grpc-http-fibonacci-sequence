package repository

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
)

type RedisRepo struct {
	cfg *config.Config
	cache *cache.Cache
}

func NewRedisRepo(cache *cache.Cache, cfg *config.Config) *RedisRepo {
	return &RedisRepo{cache: cache, cfg: cfg}
}

func (r *RedisRepo) Add(ctx context.Context, key string, value string) error {
	if err := r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL: r.cfg.Redis.FibTTL,
	}); err != nil {
		errors.Wrap(err, "redisRepo.Add.redisClient.Set")
	}
	
	return nil
}

func (r *RedisRepo) CheckFib(ctx context.Context, key string) (string, bool, error) {
	var isExist bool
	var res string
	var err error
	

	if err = r.cache.Get(ctx, key, &res); err != nil {
		return "", isExist, errors.Wrap(err, "redisRepo.Add.redisClient.Get")
	}

	if res == "" {
		return res, isExist, err
	}

	// if res, err = r.redisClient.Get(ctx, key).Result(); err == redis.Nil {
	// 	logger.Info("exist: ", isExist)
	// 	return "", isExist, nil
	// } else if err != nil {
	// 	return "", isExist, errors.Wrap(err, "redisRepo.Add.redisClient.Get")
	// }
	// logger.Info("found in cache: ", res)
	isExist = true
	return res, isExist, err
}