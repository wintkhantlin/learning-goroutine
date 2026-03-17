package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped:", ctx.Err())
			return
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go worker(ctx)
	}

	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second)
}
