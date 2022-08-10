package main

import (
	"context"
	"fhgo/gRPC/stream/person"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

type personService struct {
	person.UnimplementedSearchServiceServer
}

//普通收发
func (s *personService) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := person.PersonRes{
		Name: name + "？,这是你吗",
		Sex:  1,
	}
	return &res, nil
}

//流发普通收
func (s *personService) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "收完了"})
			break
		}
	}

	return nil
}

//普通发流收
func (s *personService) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.GetName()
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		server.Send(&person.PersonRes{Name: "拿王麻子剪刀扎" + name})
		i++
	}

	return nil
}

//流收发
func (s *personService) SearchIO(server person.SearchService_SearchIOServer) error {
	i := 0
	str := make(chan string)
	go func() {
		for {
			i++
			req, _ := server.Recv()
			if i > 10 {
				str <- "概述了"
				break
			}
			str <- req.GetName()
		}
	}()
	for {
		s := <-str
		if s == "概述了" {
			break
		}
		server.Send(&person.PersonRes{Name: s})
	}
	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":1999")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personService{})
	s.Serve(l)
}
