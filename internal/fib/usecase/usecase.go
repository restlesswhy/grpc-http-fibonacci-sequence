package usecase

import (
	"context"
	"strconv"
	// "sync"
	"time"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
	// "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

type fibUC struct {
	cfg       *config.Config
	redisRepo fib.RedisRepository
}

func NewFibUC(cfg *config.Config, redisRepo fib.RedisRepository) fib.UseCase {
	return &fibUC{
		cfg:       cfg,
		redisRepo: redisRepo,
	}
}

func (f *fibUC) GetSeq(ctx context.Context, from, to int32) (models.FibSeq, error) {
	t := time.Now()
	var fibSeq models.FibSeq
	// var wg sync.WaitGroup
	// var mutex sync.Mutex 
	fibSeq.Seq = make(map[int32]int64)

	interval := makeInterval(int(from), int(to))
	logger.Info(interval)
	for _, v := range interval {
		// wg.Add(1)
		// go func(v int, mutex *sync.Mutex) {
			// mutex.Lock()
			logger.Info("stuf with ", v)
			fibItem := f.getFib(ctx, int32(v))
			fibSeq.Seq[int32(v)] = fibItem
		// 	wg.Done()
		// 	mutex.Unlock()
		// }(v, &mutex)
		// logger.Info("done")
	}
	// wg.Wait()

	logger.Info(time.Since(t))
	return fibSeq, nil
}

func (f *fibUC) getFib(ctx context.Context, n int32) int64 {
	if n <= 1 {
		return int64(n)
	}

	nStr := strconv.Itoa(int(n))
	if res, ok, _ := f.redisRepo.CheckFib(ctx, nStr); ok {
		return res
	} 

	sum := f.getFib(ctx, n-1) + f.getFib(ctx, n-2)

	if err := f.redisRepo.Add(ctx, nStr, sum); err != nil {
		logger.Errorf("fibUC.getFib.redisRepo.Add: %s", err)
	}

	return sum
}

func makeInterval(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
