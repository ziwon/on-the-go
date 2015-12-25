package main

import (
	"fmt"
	"time"
)

const (
	OVERLOAD_SIZE       = 5
	MAX_ACCEPTABLE_SIZE = 3
	RATE_LIMIT_PER_MS   = time.Millisecond * 200
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(RATE_LIMIT_PER_MS)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, MAX_ACCEPTABLE_SIZE)
	for i := 0; i < MAX_ACCEPTABLE_SIZE; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// Add new bursty size on every RATE_LIMIT_PER_MS
		for t := range time.Tick(RATE_LIMIT_PER_MS) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests,
	// the first 3 of these will benefit from `burstyLmiter`
	burstyRequests := make(chan int, OVERLOAD_SIZE)
	for i := 1; i <= OVERLOAD_SIZE; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
