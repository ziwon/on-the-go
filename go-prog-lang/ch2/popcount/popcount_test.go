package popcount_test

import (
	"github.com/patriz/on-the-go/go-prog-lang/ch2/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCount(0xffff)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountLoop(0xffff)
	}
}

func BenchmarkPopCountShnft64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountShift64(0xffff)
	}
}

func BenchmarkPopCountClearRightmostBit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountClearRightmostBit(0xffff)
	}
}

/*
> go test -bench=.
PASS
BenchmarkPopCount-4                     300000000                5.48 ns/op
BenchmarkPopCountLoop-4                 100000000               14.6 ns/op
BenchmarkPopCountShnft64-4              100000000               18.2 ns/op
BenchmarkPopCountClearRightmostBit-4    100000000               18.7 ns/op
ok      github.com/patriz/on-the-go/go-prog-lang/ch2/popcount   7.423s
*/
