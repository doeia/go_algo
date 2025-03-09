package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// define a remote method
type Hello struct {
}

func (this Hello) SayHello(req string, res *string) error {
	*res = "hello" + req
	return nil
}

func main() {
	fmt.Println("hello server")

	// register server
	// err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}

	// register rpc
	// listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}

	// disconnect when exict
	defer listener.Close()

	for {
		fmt.Println("start connect")
		// build server
		// conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}

		// bind server
		// rpc.ServeConn(conn)
	}
}

/*
	rpc -> jsonrpc
	1. rep.Dial -> net.Dial
	2. clinet := rpc.NewClinetWithVCodec(jsonrpc.NewClient)
	3. conn.Call -> clinet.Call
*/
