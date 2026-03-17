package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan string)
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		jobs <- "job-1"
		done <- true
	}()

	fmt.Println("=== With default (non-blocking) ===")
	start := time.Now()
	for {
		select {
		case job := <-jobs:
			fmt.Println(time.Since(start), "Processing", job)
		case <-done:
			fmt.Println(time.Since(start), "All jobs done")
			break
		default:
			fmt.Println(time.Since(start), "No jobs yet, doing other work...")
			time.Sleep(500 * time.Millisecond)
		}

		if len(jobs) == 0 && len(done) == 0 && time.Since(start) > 3*time.Second {
			break
		}
	}

	fmt.Println("\n=== Without default (blocking) ===")
	go func() {
		time.Sleep(2 * time.Second)
		jobs <- "job-1"
		done <- true
	}()
	start = time.Now()
	for {
		select {
		case job := <-jobs:
			fmt.Println(time.Since(start), "Processing", job)
		case <-done:
			fmt.Println(time.Since(start), "All jobs done")
		}
	}
}
