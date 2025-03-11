package main

import (
	"context"
	"demo_client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, err: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewDemoServiceClient(conn)
	ctx, cacnel := context.WithTimeout(context.Background(), time.Second)
	defer cacnel()

	name := "Alice"
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayHello failed, err: %v", err)
		return
	}
	log.Printf("resp: %v", resp.Message)
}
