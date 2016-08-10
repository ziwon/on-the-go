package counter

import (
	"io"
	"sync"
)

type cc struct {
	count   *int64
	wrapper io.Writer
}

var ccFree = sync.Pool{
	New: func() interface{} { return new(cc) },
}

func (c *cc) doCount(b []byte) {
	*c.count = int64(len(b))
}

func (c *cc) Write(b []byte) (n int, err error) {
	c.doCount(b)
	n, err = c.wrapper.Write(b)
	return
}

func newCounter(w io.Writer) *cc {
	c := ccFree.Get().(*cc)
	c.wrapper = w
	c.count = new(int64)
	return c
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := newCounter(w)
	return c, c.count
}
