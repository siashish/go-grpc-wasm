package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/siashish/go-grpc/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.GreetingServiceServer
}

func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReponse, error) {
	return &pb.GreetingServiceReponse{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
