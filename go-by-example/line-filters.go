package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Wrapping the unbuffered `os.Stdin` with a buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text() returns the current token
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	// Check for errors during Scan()
	// End of file is expected but not reported as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
