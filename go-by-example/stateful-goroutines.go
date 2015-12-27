// sharing memory by communicating and having each piece of data owned by
// exactly 1 goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// in order to read or write our state
// other goroutines will send messages to the owning goroutine
// and receive corresponding replies by `readOp` and `writeOp`
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		// state will be owned by a single goroutine.
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
				fmt.Println("state in read:", state)
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
				fmt.Println("state in writes:", state)
			}
		}
	}()

	// starts 100 goroutines to issue reads to the state-owning goroutine
	// via `read` channel
	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// start 10 goroutines to issue writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
