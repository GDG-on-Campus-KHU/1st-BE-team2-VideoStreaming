package main

import (
	"log"
	"net"

	"github.com/kmin1231/proj_grpc/pkg/grpcserver"
)

func main() {
	// creates a listener -> waits for incoming TCP connections on port :50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// creates a new gRPC server instance -> initializes gRPC server
	s := grpcserver.NewServer()
	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
