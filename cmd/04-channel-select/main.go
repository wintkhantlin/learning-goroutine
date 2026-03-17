package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, quit <-chan struct{}) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("worker %d processing job %d\n", id, job)
			time.Sleep(time.Second)
		case <-quit:
			fmt.Printf("worker %d shutting down\n", id)
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	quit := make(chan struct{})

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, quit)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- j
	}

	time.Sleep(2 * time.Second)

	close(quit)

	time.Sleep(time.Second)
}
