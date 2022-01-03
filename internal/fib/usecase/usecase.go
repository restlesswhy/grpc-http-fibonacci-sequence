package usecase

import (
	"context"
	"math/big"
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
	fibSeq.Seq = make(map[int32]string)

	interval := makeInterval(int(from), int(to))
	logger.Info(interval)
	for _, v := range interval {
		fibItem := f.getFib(ctx, uint(v))
		fibItemStr := fibItem.String()
		fibSeq.Seq[int32(v)] = fibItemStr
	}

	logger.Info(time.Since(t))
	return fibSeq, nil
}

func (f *fibUC) getFib(ctx context.Context, n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	res := new(big.Int)
	nStr := strconv.Itoa(int(n))
	if fib, ok, _ := f.redisRepo.CheckFib(ctx, nStr); ok {
		res, _ := res.SetString(fib, 10)
		return res
	}
	
	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	n1Str := n1.String()
	if err := f.redisRepo.Add(ctx, nStr, n1Str); err != nil {
		logger.Errorf("fibUC.getFib.redisRepo.Add: %s", err)
	}

	return n1
}

func makeInterval(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
