package main

import (
	"fmt"
	"github.com/liuyibao/golang-learn/rpc/rpc03/rpclib"
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
	err = client.Call(rpclib.HelloServiceName + ".Hello", "liuyibao", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
