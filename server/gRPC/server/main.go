package main

import (
	"context"
	hello "fhgo/gRPC/protobuf"
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
	l, _ := net.Listen("tcp", ":1998")
	s := grpc.NewServer()
	hello.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
