package main

import (
	"context"
	hello "fhgo/gRPC/protobuf"
	"fhgo/service/registry"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello.Req) (res *hello.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello.Res{Message: "我是从服务端返回的gRPC的内容"}, nil
}

func main() {

	addr := "localhost:1988"
	registry.NewProducer("http://127.0.0.1:5486/_rpc_/registry/registry", "test", "hello", addr, 0)
	l, _ := net.Listen("tcp", addr)
	s := grpc.NewServer()
	hello.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
