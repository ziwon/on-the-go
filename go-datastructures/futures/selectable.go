// https://github.com/Workiva/go-datastructures/blob/master/futures/selectable.go

package futures

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrFutureCanceled = errors.New("future canceled")

type Selectable struct {
	m      sync.Mutex
	val    interface{}
	err    error
	wait   chan struct{}
	filled uint32
}

func NewSelectable() *Selectable {
	return &Selectable{}
}

func (f *Selectable) wchan() <-chan struct{} {
	f.m.Lock()
	if f.wait == nil {
		f.wait = make(chan struct{})
	}
	ch := f.wait
	f.m.Unlock()
	return ch
}

func (f *Selectable) WaitChan() <-chan struct{} {
	if atomic.LoadUint32(&f.filled) == 1 {
		return closed
	}
	return f.wchan()
}

func (f *Selectable) GetResult() (interface{}, error) {
	if atomic.LoadUint32(&f.filled) == 0 {
		<-f.wchan()
	}
	return f.val, f.err
}

func (f *Selectable) Fill(v interface{}, e error) error {
	f.m.Lock()
	if f.filled == 0 {
		f.val = v
		f.err = e
		atomic.StoreUint32(&f.filled, 1)
		w := f.wait
		f.wait = closed
		if w != nil {
			close(w)
		}
	}
	f.m.Unlock()
	return f.err
}

func (f *Selectable) SetValue(v interface{}) error {
	return f.Fill(v, nil)
}

func (f *Selectable) SetError(e error) {
	f.Fill(nil, e)
}

func (f *Selectable) Cancel() {
	f.SetError(ErrFutureCanceled)
}

var closed = make(chan struct{})

func init() {
	close(closed)
}
