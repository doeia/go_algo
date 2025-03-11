package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	X, Y int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d+%d=%d\n", args.X, args.Y, reply)

	// async call
	var reply2 int
	asyncCall := client.Go("ServiceA.Add", args, &reply2, nil)
	<-asyncCall.Done
	fmt.Printf("Arith: %d+%d=%d\n", args.X, args.Y, reply2)

}
