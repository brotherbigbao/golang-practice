package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//RPC服务端
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	log.Println("accept before")
	conn, err := listener.Accept()
	log.Println("accept after")
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	log.Println("serveConn before")
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	log.Println("serveConn After")


}

type HelloService struct {

}

func (p *HelloService) Hello(request string, reply *string) error  {
	*reply = "hello:" + request
	return nil
}
