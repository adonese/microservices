package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adonese/microservices/key/key"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentAPIClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.Request{
		TerminalID: "18000377",
		TranDateTime: "191222132700",
		STAN: 32,
		ClientID: "ACTS",
	}
	r, err := c.GetWorkingKey(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.GetWorkingKey())
}
