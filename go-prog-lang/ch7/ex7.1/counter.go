package counter

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	reader := strings.NewReader(string(p))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	err := scanner.Err()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	*c += WordCounter(count)
	return count, err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	reader := strings.NewReader(string(p))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}

	err := scanner.Err()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	*c += LineCounter(count)
	return count, err
}
