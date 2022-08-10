package main

import (
	"context"
	"fhgo/gRPC/stream/person"
	"fmt"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, _ := grpc.Dial("localhost:1999", grpc.WithInsecure())
	defer conn.Close()
	client := person.NewSearchServiceClient(conn)
	//普通收发
	//res, err := client.Search(context.Background(), &person.PersonReq{
	//	Name: "lyk",
	//	Sex:  1,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	//流发普通收
	//c, _ := client.SearchIn(context.Background())
	//i := 0
	//for {
	//	if i > 10 {
	//		res, _ := c.CloseAndRecv()
	//		fmt.Println(res)
	//		break
	//	}
	//	c.Send(&person.PersonReq{Name: "进去了"})
	//	time.Sleep(1 * time.Second)
	//	i++
	//}

	//普通发流收
	//c, _ := client.SearchOut(context.Background(), &person.PersonReq{Name: "王大麻子"})
	//for {
	//	req, err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(req)
	//}

	//流收发
	c, _ := client.SearchIO(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			err := c.Send(&person.PersonReq{Name: "客户端发的"})
			if err != nil {
				fmt.Print("收的err：")
				fmt.Println(err)
				break
			}
		}
		wg.Done()
	}()
	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				fmt.Print("发的err：")
				fmt.Println(err)
				break
			}
			fmt.Println(res)
		}
		wg.Done()
	}()

	wg.Wait()
}
