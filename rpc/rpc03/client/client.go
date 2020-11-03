package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//RPC客户端
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "liuyibao", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
