// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func check(u string, checked chan<- bool) {
	time.Sleep(4 * time.Second)
	checked <- true
}

func Reachable(urls []string) bool {
	ch := make(chan bool, 1)
	for _, url := range urls {
		go func(u string) {
			checked := make(chan bool)
			go check(u, checked)
			select {
			case ret := <-checked:
				ch <- ret
			case <-time.After(5 * time.Second):
				ch <- false
			}
		}(url)
	}
	return <-ch
}

func main() {
	fmt.Println(Reachable([]string{"url1"}))
}
