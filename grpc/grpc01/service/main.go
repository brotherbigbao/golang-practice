package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	_ "log"
	"net"
	_ "net"
	_ "google.golang.org/grpc"
	"github.com/liuyibao/golang-learn/grpc/grpc01/helloservice"
)

type HelloServiceImpl struct {

}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *helloservice.String) (*helloservice.String, error) {
	reply := &helloservice.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	helloservice.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
