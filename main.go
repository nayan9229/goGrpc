package main

import (
	"log"
	"net"

	"github.com/nayan9229/goGrpc/api/proto"
	"github.com/nayan9229/goGrpc/pkg/adder"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// Register gRPC server
	// api.RegisterAdderServer(s, srv)
	api.RegisterAdderServer(s, srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
