package main

import (
	"fmt"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	w1 := "世界"
	w2 := "\xe4\xb8\x96\xe7\x95\x8c" // not a legal rune liternal
	w3 := "\u4e16\u754c"
	w4 := "\U00004e16\U0000754c"

	fmt.Println(w1, w2, w3, w4)

	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s = "푸로구라무"
	fmt.Printf("%x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(r))
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(123457688))

}
