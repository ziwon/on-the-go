package main

import (
	"fmt"
	"time"
)

func worker(ticker *time.Ticker, quit chan bool) {
	for {
		select {
		case <-ticker.C:
			pings := make(chan string, 1)
			pongs := make(chan string, 1)
			ping(pings, "passed message")
			pong(pings, pongs)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func ping(pings chan<- string, msg string) {
	fmt.Println("ping >> ", msg)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("pong << ", msg)
	pongs <- msg
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan bool)
	go worker(ticker, quit)
	time.Sleep(5 * time.Second)
	quit <- true
}
