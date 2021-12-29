package grpcdel

import (
	"context"

	pb "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/proto"
)

type FibMicroservice struct {
	pb.UnimplementedFiboSequenceServiceServer

}

func NewFibMicroservice() *FibMicroservice {
	return &FibMicroservice{}
}

func (f *FibMicroservice) Get(context.Context, *pb.FiboRequest) (*pb.FiboResponse, error) {
	return nil, nil
}