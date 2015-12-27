// Package main provides ...
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var ops int64 = 0

	// start 100 goroutines
	for i := 0; i < 100; i++ {
		go func(i int) {
			total := 0
			for {
				key := rand.Intn(5)

				mutex.Lock()
				total += state[key]
				fmt.Println("i:", i, "state:", state)
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// ensure that this goroutine doesn't starve the scheduler
				// by yielding `runtime.Gosched()`
				runtime.Gosched()
			}
		}(i)
	}

	// start 10 goroutines to write random value and increment the `ops` count
	for w := 0; w < 10; w++ {
		go func(w int) {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)

				mutex.Lock()
				state[key] = val
				fmt.Println("w:", w, "state:", state)
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}(w)
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("last ops:", opsFinal)

	mutex.Lock()
	fmt.Println("last state:", state)
	mutex.Unlock()
}
