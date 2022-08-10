package main

import (
	"context"
	hello "fhgo/gRPC/protobuf"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:1998", grpc.WithInsecure())
	defer conn.Close()
	client := hello.NewHelloGRPCClient(conn)
	res, _ := client.SayHi(context.Background(), &hello.Req{Message: "你好，我是客户端"})
	fmt.Println(res.GetMessage())
}
