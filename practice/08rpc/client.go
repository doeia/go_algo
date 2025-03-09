package client

import (
	"fmt"
	"net/rpc"
)

func client() {
	// conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}

	defer conn.Close()

	// call remote func
	var reply string
	// err2 := conn.Call("hello.SayHello", "client", &reply)

	if err2 != nil {
		fmt.Println(err2)
	}

	// get remote info
	fmt.Println(reply)
}
