package main

import (
	"fmt"
)

// Given the following program
// There are in three possible results with race conditions
//
// nothing printed:
// 1:
// 0:
func main() {
	var data int
	go func() { // <1>
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
