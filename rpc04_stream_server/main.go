package main

import (
	"demo_server/pb"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// this is the server struct that implements the interface
	pb.UnimplementedDemoServiceServer
}

// func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.DemoService_LotsOfRepliesServer) error {
// 	words := []string{"Hellor", "你好", "안녕하세요", "こんにちは"}

// 	for _, word := range words {
// 		data := &pb.HelloResponse{
// 			Message: word + " " + in.GetName(),
// 		}
// 		if err := stream.Send(data); err != nil {
// 			return err
// 		}
// 	}
// 	return nil

// }

func (s *server) LotsOfGreetings(stream pb.DemoService_LotsOfGreetingsServer) error {
	reply := "hello"
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{Message: reply})
		}
		if err != nil {
			return err
		}
		reply = "hello " + in.GetName()
	}
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
