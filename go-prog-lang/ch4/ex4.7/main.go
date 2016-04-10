package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 세상아"

	buf := make([]byte, utf8.RuneCountInString(s)+1)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		b := make([]byte, 3)
		_ = utf8.EncodeRune(b, r)
		buf = append(buf, b...)
		i += size
	}

	fmt.Println(string(buf)) //Hello, 세상아
	reverse(buf)
	fmt.Println(string(buf)) //<84><95>쁃츄 ,olleH
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
