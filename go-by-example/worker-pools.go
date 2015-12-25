package main

import (
	"fmt"
	"time"
)

const (
	WORKER = 10
	JOBS   = 100
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing jobs", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= WORKER; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= JOBS; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= JOBS; a++ {
		<-results
	}
	close(results)
}
