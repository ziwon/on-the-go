package main

import (
	"fmt"
)

var global *int

func f() *int {
	var x int
	x = 1
	global = &x
	return global
}

func g() *int {
	y := new(int)
	*y = 1
	return y
}

func main() {
	fmt.Println(f() == g())
}
