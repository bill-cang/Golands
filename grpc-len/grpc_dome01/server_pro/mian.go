package main

import (
	"../helloworld"
	"context"
	"fmt"
	"github.com/lunny/log"
	"google.golang.org/grpc"
	"net"
	"utils"
)

const (
	port = ":8081"
)

//1、定义结构体用来实现server接口
type server struct {

}

func (s *server) SayHello(ctx context.Context, req *halloWorld.HelloRequest) (*halloWorld.HelloReply, error) {
	str := fmt.Sprintf("hello %s , this is server SayHello! " , req.Name)
	var reply = halloWorld.HelloReply{
		Message: str,
	}
	return &reply, nil
	//status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func main() {
	//监听套接字
	listener, e := net.Listen("tcp", port)
	utils.HandleError(e, "net.Listen")

	s := grpc.NewServer()
	//注册到grpc server
	halloWorld.RegisterGreeterServer(s, &server{})

	//s.Serve(listener) 无限循环
	if err := s.Serve(listener); err != nil {
		log.Fatal("fail to server")
	}

}
