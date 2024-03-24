package main

import (
	"context"
	"io"
	"log"

	pb "github.com/turanheydarli/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Print("doGreetManyTimes invoked")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Turan",
	})

	if err != nil {
		log.Printf("Could not greet %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error while greet %v", err)
		}

		log.Printf("Message %s\n", message.Result)
	}
}
