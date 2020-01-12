package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/adonese/microservices/raterpc/rate"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetSDGRate(ctx context.Context, in *pb.Empty) (*pb.SDGRate, error) {

	rate := getRate()
	return &pb.SDGRate{Message: rate}, nil
}

func getRate() float32 {
	a := extract("https://www.price-today.com/currency-prices-sudan/")
	fmt.Printf("The values are: %v\n", a)
	_, r := getUSD(a)
	fmt.Printf("The rate currently is: %v\n", r)

	f, _ := strconv.ParseFloat(r, 32)

	return float32(f)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	var t server
	pb.RegisterRaterServer(s, &t)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
