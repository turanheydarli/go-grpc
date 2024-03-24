package main

import (
	"context"
	"log"

	pb "github.com/turanheydarli/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Print("doGreet invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Turan",
	})

	if err != nil {
		log.Printf("Could not greet %v", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
