package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

// both goroutines wait infinitely on each other
// deadlock prevention works by preventing one of the four Coffman conditions from occurring.


// mutual exclusion: a concurrent process holds exclusive rights to a resource at any one time.
// wait for condition: a concurrent process must simultaneously hold a resource and be waiting for an additional resource.
// no preemption: a resource held by a concurrent process can only be released by that process.
// circular wait: a concurrent process(P1) must be waiting on a chain of other concurrent process (P2), which are in turn waiting on it (P1)
func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		defer v1.mu.Unlock()
		defer v2.mu.Unlock()
		v1.mu.Lock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
