package main

import (
	"context"
	"log"

	// pb "github.com/restlesswhy/grpc/shortener-client/proto"

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
		From: 3,
		To: 10,
	})
	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	// ress, err := client.Get(context.Background(), &pb.UGRequest{
	// 	ShortUrl: "WOEAi4MH23",
	// })
	// if err != nil {
	// 	log.Fatalf("could not get answer: %v", err)
	// }


	log.Println("Sequence:", resp.Result)
	// log.Println("Long url:", ress.LongUrl)
}
