package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q[2])

	s := [...]int{1, 2, 3}
	fmt.Printf("%T\n", s)

	t := [3]int{1, 2, 3}
	fmt.Println(t)

	//t = [4]int{1, 2, 3, 4}
	//fmt.Println(t)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "L", RMB: "Y"} // index as enum value
	fmt.Println(RMB, symbol[RMB])

	z := [...]int{99: -1}
	fmt.Println(z) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 -1]

	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}

	fmt.Println(a1 == b1, a1 == c1, b1 == c1)
	// d1 := [3]int{1, 2}
	// fmt.Println(a1 == d1) // mismatched type [2]int and [3]int

}
