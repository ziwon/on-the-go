package main

import (
	"fmt"
	"os"
)

const Space = " "

func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		if i > 1 {
			s += Space
		}
		s += os.Args[i]
	}
	fmt.Fprintln(os.Stdout, s)
}
