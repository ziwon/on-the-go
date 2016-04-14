package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func Walk(t *tree, ch chan int) {
	if t.left != nil {
		Walk(t.left, ch)
	}

	if t.right != nil {
		Walk(t.right, ch)
	}

	if t != nil {
		ch <- t.value
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	ints := make([]int, 10)
	for i, v := range ints {
		v = rand.Intn(1000)
		fmt.Printf("init: %d\n", v)
		ints[i] = v
	}

	ch := make(chan int)
	go func() {
		Walk(Sort(ints), ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("sort: %d\n", v)
	}

	/*
		init: 81
		init: 887
		init: 847
		init: 59
		init: 81
		init: 318
		init: 425
		init: 540
		init: 456
		init: 300

		sort: 59
		sort: 300
		sort: 456
		sort: 540
		sort: 425
		sort: 318
		sort: 81
		sort: 847
		sort: 887
		sort: 81
		╭
	*/

}
