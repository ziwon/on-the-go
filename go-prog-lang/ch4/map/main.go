package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	/*
		ages := map[string]int {
			"alice": 31,
			"charlie": 34,
		}
	*/

	//fmt.Println(ages["alice"])
	//delete(ages, "alice")

	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages["bob"])

	// map element is not a variable, cannot take its address
	// _ = &ages["bob"]

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["no_key"]; !ok {
		fmt.Println(age)
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
