package usecase

import (
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
)

type fibUC struct {
	cfg *config.Config
	redisRepo fib.RedisRepository
}

func NewFibUC(cfg *config.Config, redisRepo fib.RedisRepository) fib.UseCase {
	return &fibUC{
		cfg: cfg,
		redisRepo: redisRepo,
	}
}

func (f *fibUC) Get(x, y int32) {

}