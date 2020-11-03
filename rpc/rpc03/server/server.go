package main

import (
	"github.com/liuyibao/golang-learn/rpc/rpc03/rpclib"
	"log"
	"net"
	"net/rpc"
)

//RPC服务端
func main() {
	//rpc.RegisterName("HelloService", new(HelloService))
	rpclib.RegisterHelloService(new (HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

type HelloService struct {

}

func (p *HelloService) Hello(request string, reply *string) error  {
	*reply = "hello:" + request
	return nil
}
