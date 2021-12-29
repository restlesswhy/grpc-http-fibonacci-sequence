package fib

type RedisRepository interface {
	CheckFib(x int32) (bool, error)	
	Add(x int32, y int64) error
}

