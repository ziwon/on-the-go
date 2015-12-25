// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func download(u string, ch chan<- bool) {
	time.Sleep(4 * time.Second)
	ch <- true
}

func Reachable(urls []string) bool {
	ch := make(chan bool, len(urls))
	for _, url := range urls {
		go download(url, ch)
	}
	time.AfterFunc(time.Second, func() { ch <- false })
	return <-ch
}

func main() {
	fmt.Println(Reachable([]string{"url1", "url2"}))
}
