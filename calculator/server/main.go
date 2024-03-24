package main

import (
	"log"
	"net"

	pb "github.com/turanheydarli/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", addr)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", addr)
	}
}
