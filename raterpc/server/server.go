package main

import (
	"context"
	"errors"
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

func (s *server) GetDonations(ctx context.Context, in *pb.DonationURL) (*pb.TotalDonations, error) {
	var e ebs
	ok, data := e.getOnline(in.GetUrl())
	if !ok {
		return nil, errors.New("couldn't get exact data from ebs")
	}
	log.Printf("The data is: %s, %s", data[0], data[1])
	number, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("unable to convert data: %v", err)
	}
	amount, err := strconv.Atoi(data[1])

	if err != nil {
		return nil, fmt.Errorf("unable to convert data: %v", err)
	}

	return &pb.TotalDonations{TotalAmount: float32(number), NumberTransactions: int32(amount)}, nil
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
