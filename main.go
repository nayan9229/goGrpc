package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/nayan9229/goGrpc/api/proto"
	"github.com/nayan9229/goGrpc/pkg/adder"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// Client function which will request to server every 5 sec and print the result
	go func() {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":8080", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()
		ac := api.NewAdderClient(conn)
		count := int32(0)
		for {
			time.Sleep(5 * time.Second)
			response, err := ac.Add(context.Background(), &api.AddRequest{
				X: 10 * count,
				Y: 20 * count,
			})
			count++
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			log.Printf("Response from server: %d", response.Result)
		}
	}()

	// Register gRPC server
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
