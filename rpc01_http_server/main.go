package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	server := new(ServiceA)
	rpc.Register(server)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(listener, nil)
}
