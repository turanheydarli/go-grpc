package main

import (
	"context"
	"io"
	"log"

	pb "github.com/turanheydarli/grpc-go/calculator/proto"
)

func doPrimes(c pb.SumServiceClient) {
	log.Print("doPrimes invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Printf("error %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error while stream %v", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}

}
