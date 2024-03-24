package main

import (
	"context"
	"log"

	pb "github.com/turanheydarli/grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked %v", in)

	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
