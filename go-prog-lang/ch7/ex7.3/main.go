package main

import (
	_ "bytes"
	"fmt"
	"math/rand"
	_ "sync"
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

func (c *tree) String() string {
	ch := make(chan int)

	go func() {
		Walk(c, ch)
		close(ch)
	}()

	var acc string
	for c := range ch {
		acc += fmt.Sprintf("%d ", c)
	}

	return acc
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	ints := make([]int, 10)
	for i, v := range ints {
		v = randInt(1, 100)
		ints[i] = v
	}

	tree := Sort(ints)
	fmt.Println(tree)
}
