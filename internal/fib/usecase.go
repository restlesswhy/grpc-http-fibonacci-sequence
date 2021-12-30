package fib

type UseCase interface {
	GetSeq(x, y int32) (map[int32]int64, error)
}