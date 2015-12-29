package main

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger(resume chan int) int {
	return <-resume
}

func main() {
	resume := integers()
	fmt.Println("generateInteger() -> ", generateInteger())
	fmt.Println("generateInteger() -> ", generateInteger())
	fmt.Println("generateInteger() -> ", generateInteger())
}
