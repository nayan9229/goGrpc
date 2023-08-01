package main

import (
	// "context"
	"fmt"
	"log"
	"net"
	// "time"

	"github.com/nayan9229/goGrpc/chat"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := chat.NewChatServiceClient(conn)

		for {
			time.Sleep(3 * time.Second)
			response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			log.Printf("Response from server: %s", response.Body)
		}
	}()

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
