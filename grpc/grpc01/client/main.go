package main

import (
	"context"
	"fmt"
	"github.com/liuyibao/golang-learn/grpc/grpc01/helloservice"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := helloservice.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &helloservice.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
