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

func (f *FibMicroservice) Get(ctx context.Context, in *pb.FiboRequest) (*pb.FiboResponse, error) {

	y, _ := f.fibUC.GetSeq(ctx, in.From, in.To)

	return &pb.FiboResponse{
		Result: y.Seq,
	}, nil
}