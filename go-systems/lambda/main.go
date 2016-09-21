package main

import "fmt"

func make2(x int) []func(int) int {
	return []func(int) int{
		func(a int) int { x += a; return x },
		func(b int) int { x += b; return x },
	}
}

func add(x int) func() int {
	return func() int { x += 1; return x }
}

func add2(x int) func(int) func(int) int {
	return func(a int) func(int) int {
		x += a
		return func(b int) int {
			x += b
			return x
		}
	}
}

func main() {
	f := make2(1)
	for i := 1; i <= 3; i++ {
		fmt.Println(f[0](i))
		fmt.Println(f[1](2 * i))
	}

	addFrom2 := add(2)
	addFrom3 := add(3)

	fmt.Println(addFrom2()) // 3
	fmt.Println(addFrom2()) // 4
	fmt.Println(addFrom3()) // 4
	fmt.Println(addFrom3()) // 5

	f1 := add2(3)           // 3
	f2 := f1(4)             // +4
	fmt.Println(f1(10)(11)) // 28: 3 + 4 + 10 + 11
	fmt.Println(f2(4))      // 32: 28 + 4
	fmt.Println(f2(5))      // 37: 32 + 5
}
