// http://stackoverflow.com/questions/12265813
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const SizePerThread = 1000000

type Queue struct {
	records string
	count   int64
}

func (q *Queue) push(record chan interface{}) {
	record <- time.Now()

	newcount := atomic.AddInt64(&q.count, -1)
	log.Printf("Push: %d\n", newcount)
}

func (q *Queue) pop(record chan interface{}) {
	<-record

	newcount := atomic.AddInt64(&q.count, -1)
	log.Printf("Pop: %d\n", newcount)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup

	runtime.GOMAXPROCS(runtime.NumCPU())

	record := make(chan interface{}, 10000000)
	queue := new(Queue)

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < SizePerThread; j++ {
				queue.push(record)
			}
		}()

		go func() {
			defer wg.Done()

			for i := 0; i < SizePerThread; i++ {
				queue.pop(record)
			}
		}()
	}

	wg.Wait()
}
