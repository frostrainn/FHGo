package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		go watch(ctx, 1)
		go watch(ctx, 2)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("watch %d %s\n", 0, ctx.Err())
	}

	fmt.Println("finished")
}

func watch(ctx context.Context, flag int) {
	fmt.Printf("doing something flag:%d\n", flag)
	time.Sleep(5 * time.Second)
	fmt.Println("finished flag:", flag)
}
