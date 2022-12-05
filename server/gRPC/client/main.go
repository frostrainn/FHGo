package main

import (
	"context"
	hello "fhgo/gRPC/protobuf"
	"fhgo/service/registry"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"time"
)

func main() {

	consumer := registry.NewConsumer("http://127.0.0.1:5486/_rpc_/registry/registry",
		"gRPC", "hello")
	resolver.Register(consumer)
	conn, _ := grpc.Dial("zhubi://hello", grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy":"%s"}`, roundrobin.Name)), grpc.WithInsecure())
	defer conn.Close()
	client := hello.NewHelloGRPCClient(conn)
	for {
		res, _ := client.SayHi(context.Background(), &hello.Req{Message: "我是客户端的内容"})
		fmt.Println(res.GetMessage())
		time.Sleep(time.Second * 20)
	}
}
