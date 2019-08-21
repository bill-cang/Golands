package main

import (
	"../helloworld"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"utils"
)

const address = "127.0.0.1:8081"

func main() {
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	utils.HandleError(e, "grpc.Dial")
	defer conn.Close()

	client := halloWorld.NewGreeterClient(conn)
	res, e := client.SayHello(context.Background(), &halloWorld.HelloRequest{Name: "ckx"})
	utils.HandleError(e, "halloWorld.NewGreeterClient.SayHello")

	fmt.Println(res)
}
