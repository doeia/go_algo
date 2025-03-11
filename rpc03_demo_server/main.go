package main

import (
	"context"
	"demo_server/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// this is the server struct that implements the interface
	pb.UnimplementedDemoServiceServer
}

// the method to be called
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "Hello " + in.GetName()
	return &pb.HelloResponse{Message: reply}, nil
}

func main() {
	// listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	// register server
	s := grpc.NewServer()
	pb.RegisterDemoServiceServer(s, &server{})
	// serve
	err = s.Serve(l)

	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
