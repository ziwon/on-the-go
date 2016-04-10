package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a, 2)) // [2 3 4 5 0 1]
	fmt.Println(rotate(a, 3)) // [3 4 5 0 1 2]
}

func rotate(a []int, i int) []int {
	a, b := a[:i], a[i:]
	return append(b, a...)

}
