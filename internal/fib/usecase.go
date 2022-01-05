//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock
package fib

import (
	"context"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
)

type UseCase interface {
	GetSeq(ctx context.Context, from, to int32) (models.FibSeq, error)
}