package popcount_test

import (
	"github.com/ziwon/on-the-go/go-prog-lang/ch2/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCount(0x1f)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountLoop(0x1f)
	}
}

func BenchmarkPopCountShnft64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountShift64(0x1f)
	}
}

func BenchmarkPopCountClearRightmostBit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountClearRightmostBit(0x1f)
	}
}

func BenchmarkPopCountMultipleBy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountMultipleBy(0x1f)
	}
}

func BenchmarkPopCountAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcount.PopCountAdd(0x1f)
	}
}

/*
> go test -bench=.
PASS
BenchmarkPopCount-4                     300000000                5.48 ns/op
BenchmarkPopCountLoop-4                 100000000               14.6 ns/op
BenchmarkPopCountShnft64-4              100000000               18.2 ns/op
BenchmarkPopCountClearRightmostBit-4    100000000               18.7 ns/op
BenchmarkPopCountMultipleBy-4           1000000000               2.33 ns/op
BenchmarkPopCountAdd-4                  300000000                4.05 ns/op
ok      github.com/ziwon/on-the-go/go-prog-lang/ch2/popcount   7.423s
*/
