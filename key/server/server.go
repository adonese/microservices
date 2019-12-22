package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"

	pb "github.com/adonese/microservices/key/key"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetWorkingKey(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	req := EbsRequest{
		TranDateTime: in.TranDateTime,
		TerminalID:   in.TerminalID,
		ClientID:     in.ClientID,
		STAN:         int(in.STAN),
	}

	url := "https://beta.soluspay.net/api/workingKey"
	b, err := json.Marshal(&req)
	if err != nil {
		log.Printf("The error is: %v", err)
		return nil, errors.New("it doesn't work")
	}
	result, err := request(b, url)
	if err != nil {
		return nil, err
	}
	log.Printf("the response is: %v", result)

	return &pb.Response{
		WorkingKey:      result.Response.WorkingKey,
		ResponseCode:    int32(result.Response.ResponseCode),
		ResponseMessage: result.Response.ResponseMessage,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	var t server
	pb.RegisterWorkingKeyServer(s, &t)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
