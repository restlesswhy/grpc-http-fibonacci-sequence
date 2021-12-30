package grpcdel

import (
	"context"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	pb "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/proto"
)

type FibMicroservice struct {
	pb.UnimplementedFiboSequenceServiceServer
	fibUC fib.UseCase
}

func NewFibMicroservice(fibUC fib.UseCase) *FibMicroservice {
	return &FibMicroservice{fibUC: fibUC}
}

func (f *FibMicroservice) Get(context.Context, *pb.FiboRequest) (*pb.FiboResponse, error) {
	return nil, nil
}