package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	var s []int

	if s = nil; len(s) == 0 && s == nil {
		fmt.Println("empty & nil")
	}

	if s = []int(nil); len(s) == 0 && s == nil {
		fmt.Println("empty & nil")
	}

	if s = []int{}; len(s) == 0 && s != nil {
		fmt.Println("empty & not nil")
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
