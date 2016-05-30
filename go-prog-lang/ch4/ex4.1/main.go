package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/ziwon/on-the-go/go-prog-lang/ch2/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(PopCountDiff(c1, c2)) // 125
	c3 := [...]byte{'a'}
	c4 := [...]byte{'b'}
	fmt.Println(PopCountDiff1(c3, c4)) // 2
}

func PopCountDiff(a, b [32]byte) int {
	var sum int
	for i := range a {
		diff := uint64(a[i] ^ b[i])
		sum += popcount.PopCount(diff)
	}
	return sum
}

func PopCountDiff1(a, b [1]byte) int {
	fmt.Printf("> %08b\n> %08b\n", a[0], b[0])
	var sum int
	for i := range a {
		diff := uint64(a[i] ^ b[i])
		sum += popcount.PopCount(diff)
	}
	return sum
}
