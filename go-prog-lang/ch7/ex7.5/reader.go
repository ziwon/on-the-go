package reader

import (
	"errors"
	"io"
)

type LimitReader struct {
	r io.Reader
	n int64
}

func (r *LimitReader) Read(p []byte) (n int, err error) {
	if r.n < 0 {
		return 0, errors.New("strings.LimitReader.Read: negative number")
	}
	if r.n >= int64(len(p)) {
		return 0, io.EOF
	}
	return copy(p, p[:r.n]), io.EOF
}

// ex7.5 Change function name to NewLimitReader from LimitReader to use it for struct type
func NewLimitReader(r io.Reader, n int64) io.Reader {
	return &LimitReader{r, n}
}
