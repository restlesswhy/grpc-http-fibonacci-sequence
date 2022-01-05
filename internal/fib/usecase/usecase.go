package usecase

import (
	"context"
	"math/big"
	"strconv"

	"time"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

type FibUC struct {
	cfg       *config.Config
	redisRepo fib.RedisRepository
}

func NewFibUC(cfg *config.Config, redisRepo fib.RedisRepository) *FibUC {
	return &FibUC{
		cfg:       cfg,
		redisRepo: redisRepo,
	}
}

// GetSeq главная функция бизнес логики, создает заданный слайс
func (f *FibUC) GetSeq(ctx context.Context, from, to int32) (models.FibSeq, error) {
	t := time.Now()
	var fibSeq models.FibSeq
	fibSeq.Seq = make(map[int32]string)

	// создаем интервал
	interval := makeInterval(int(from), int(to))
	// logger.Info(interval)
	// выполняем бизнес логику по каждому элементу интервала
	for _, v := range interval {
		fibItem := f.getFib(ctx, uint(v), f.cfg.Redis.Caching)
		fibItemStr := fibItem.String()
		fibSeq.Seq[int32(v)] = fibItemStr
	}

	logger.Info(time.Since(t))
	return fibSeq, nil
}

// getFib вычисляет число фиобанччи, работает с базой данных
func (f *FibUC) getFib(ctx context.Context, n uint, caching bool) *big.Int {

	
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	// проверяем наличие числа в базе, если отсутствует то вычисляем его и добаляем в базу
	if caching {
		res := new(big.Int)
		nStr := strconv.Itoa(int(n))

		if fib, ok, _ := f.redisRepo.CheckFib(ctx, nStr); ok {
			res, _ := res.SetString(fib, 10)
			return res
		}
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	if caching {
		resStr := n1.String()
		keyStr := strconv.Itoa(int(n))
		
		// logger.Info("create: ", resStr)
		if err := f.redisRepo.Add(ctx, keyStr, resStr); err != nil {
			logger.Errorf("fibUC.getFib.redisRepo.Add: %s", err)
		}
	}

	return n1
}

// makeInterval создает интервал по заданным параметрам
func makeInterval(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
