package main

import (
	"context"
	"demo_client/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runLotsOfReplies(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("LotsOfReplies failed, err: %v", err)
		return
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv failed, err: %v", err)
			return
		}
		log.Printf("resp: %v", resp.Message)
	}
}

func runLotsOfGreetings(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("LotsOfGreetings failed, err: %v", err)
		return
	}
	names := []string{"Alice", "Bob", "Cindy"}
	for _, name := range names {
		stream.Send(&pb.HelloRequest{Name: name})
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("stream.CloseAndRecv failed, err: %v", err)
		return
	}
	log.Printf("resp: %v", resp.Message)

}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, err: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewDemoServiceClient(conn)
	// runLotsOfReplies(c)
	runLotsOfGreetings(c)

}
