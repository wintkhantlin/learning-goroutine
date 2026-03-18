package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Input int
}

type Result struct {
	JobID  int
	Output int
}

func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // signal this goroutine is done

	for job := range jobs { // goroutine blocks here waiting for jobs
		time.Sleep(10 * time.Millisecond)

		results <- Result{ // send result to another goroutine via channel
			JobID:  job.ID,
			Output: job.Input * 2,
		}
	}
}

func main() {
	workerSize := 10
	numJobs := 10000

	jobs := make(chan Job, 100)       // shared channel between producer and workers
	results := make(chan Result, 100) // shared channel between workers and consumer

	var wg sync.WaitGroup

	for w := 1; w <= workerSize; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg) // start worker goroutines
	}

	go func() { // producer goroutine
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j, Input: j} // send jobs to workers
		}
		close(jobs) // notify workers no more jobs
	}()

	go func() { // closer goroutine
		wg.Wait()      // wait for all worker goroutines to finish
		close(results) // then close results channel
	}()

	for result := range results { // consumer loop (main goroutine)
		fmt.Println(result.Output)
	}
}
