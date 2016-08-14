package reader

import (
	"io"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	l := 5
	data := "Hello World"

	buf := make([]byte, 128)

	r := NewLimitReader(strings.NewReader(data), int64(l))
	n, err := r.Read(buf)
	if data[:n] != data[:l] {
		t.Error("Error reading at number of bytes:", l)
	}
	if err != io.EOF {
		t.Error(err)
	}
}
