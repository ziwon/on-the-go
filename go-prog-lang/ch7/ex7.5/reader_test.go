package reader

import (
	"fmt"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	offset := 3
	data := "Hello World"

	buf := make([]byte, 512)
	r := NewLimitReader(strings.NewReader(data), int64(len(data)-offset))
	n, _ := r.Read(buf)

	if data[:n] != string(buf[:n]) {
		t.Error("Ooops...")
	} else {
		fmt.Println(data[:n])
	}
}
