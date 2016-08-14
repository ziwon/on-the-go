package reader

import (
	"io"
)

type Reader struct {
	s string
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return copy(p, []byte(r.s)), io.EOF
}

func NewReader(s string) *Reader {
	return &Reader{s}
}
