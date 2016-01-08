package main

import (
	"fmt"
	"os"
)

var notFound = os.Getenv("NOT_FOUND")

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

func test() {
	fmt.Println("test called")
	if notFound == "" {
		panic("no value for $NOT_FOUND")
	}
}

func main() {
	if x := throwsPanic(test); x {
		fmt.Println("recovered")
	}

	test()
}
