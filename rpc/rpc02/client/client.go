package main

import (
	"fmt"
	"github.com/liuyibao/golang-learn/rpc/rpc02/rpclib"
	"log"
)

//RPC客户端
func main() {
	client, err := rpclib.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
