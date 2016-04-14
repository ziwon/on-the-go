package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three", "", "one"}
	fmt.Printf("%q\n", dup(data))
}

func dup(strings []string) []string {
	out := []string{}
	m := make(map[string]struct{})
	for _, s := range strings {
		if _, ok := m[s]; !ok {
			out = append(out, s)
			m[s] = struct{}{}
		}
	}
	return out
}
