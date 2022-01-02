package grpcdel

import (
	"context"
	"errors"

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

func (f *FibMicroservice) Get(ctx context.Context, in *pb.FiboRequest) (*pb.FiboResponse, error) {
	y, err := f.fibUC.GetSeq(ctx, in.From, in.To)
	if err != nil {
		return nil, err
	}

	if in.From > in.To || in.From < 0 || in.To < 0 {
		return nil, errors.New("not correct input")
	}

	return &pb.FiboResponse{
		Result: y.Seq,
	}, nil
}