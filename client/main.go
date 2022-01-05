package main

import (
	"context"
	"log"

	pb "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewFiboSequenceServiceClient(conn)
	
	resp, err := client.Get(context.Background(), &pb.FiboRequest{
		From: 1,
		To: 600,
	})
	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	log.Println("Sequence:", resp.Result)
}
