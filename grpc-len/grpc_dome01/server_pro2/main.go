package main

import (
	"../helloworld"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"utils"
)

const port = ":8081"

type server2 struct {
	halloWorld.UnimplementedGreeterServer
}

func (s *server2) SayHello(ctx context.Context, req *halloWorld.HelloRequest) (*halloWorld.HelloReply, error) {
	str := fmt.Sprintf("hello %s ,this is server2 SayHello", req.Name)
	reply := &halloWorld.HelloReply{
		Message: str,
	}
	return reply, nil
}

func main() {
	listener, e := net.Listen("tcp", port)
	utils.HandleError(e, "net.Listen")

	s := grpc.NewServer()
	//注册到grpc server
	halloWorld.RegisterGreeterServer(s, &server2{})

	if err := s.Serve(listener); err != nil {
		log.Fatal("fail to server")
	}

}
