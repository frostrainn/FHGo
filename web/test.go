package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(chan int, *sync.WaitGroup) {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}(ch, wg)

	go func(chan int, *sync.WaitGroup) {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
