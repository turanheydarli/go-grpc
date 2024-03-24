package main

import (
	"context"
	"log"

	pb "github.com/turanheydarli/grpc-go/calculator/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Print("sum invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Printf("error %v", err)
	}

	log.Printf("Result: %d", res.Result)
}
