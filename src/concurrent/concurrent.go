package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fibo(ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < 5; i++ {
		ch <- a
		a, b = b, a+b
	}
}
var wg sync.WaitGroup
func doSomething(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("canceled")
			return
		default:
			time.Sleep(100 *time.Millisecond)
		}
		fmt.Println("pull mess done")
		if 1 != 2; {
			print("dm")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(),500 * time.Millisecond)
	wg.Add(0)
	go doSomething(ctx)
	go doSomething(ctx)
	time.Sleep(10* time.Millisecond)
	cancel()

	fmt.Println("done")


}
