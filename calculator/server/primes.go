package main

import (
	"log"

	pb "github.com/turanheydarli/grpc-go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.SumService_PrimesServer) error {
	log.Printf("Primes function invoked %v", in)

	number := in.Number
	divisior := int64(2)

	for number > 1 {
		if number%divisior == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisior,
			})

			number /= divisior
		} else {
			divisior++
		}
	}

	return nil
}
