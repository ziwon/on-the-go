package main

import (
	"fmt"
	"strings"
)

func main() {
	str :=
		`Now led tedious shy lasting females off.
		ladyship so. Not$foo attention say frankness intention out dashwoods now curiosity.
		Stronger ecstatic as no judgment daughter speedily thoughts.
		Worse downs nor might she court did nay $foo forth these.`

	fmt.Println(expand(str, func(s string) string { return "" }))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
