package usecase

import (
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
)

func init() {
	cache = make(map[int32]int64)
}

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

func (f *fibUC) GetSeq(x, y int32) (map[int32]int64, error) {
	res := make(map[int32]int64)
	interval := makeInterval(int(x), int(y))
	for _, v := range interval {
		f := getFib(int32(v))
		res[int32(v)] = f
	}

	return res, nil
}

var (
	cache map[int32]int64
)


func getFib(n int32) int64 {
	if n <= 1 {
		return int64(n)
	}

	if r, ok := cache[n]; ok {
		return r
	}

	sum := getFib(n-1) + getFib(n-2)
	cache[n] = sum

	
	return sum
}

func makeInterval(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}