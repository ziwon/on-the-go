// https://github.com/Workiva/go-datastructures/blob/master/futures/futures.go
/*
Copyright 2014 Workiva, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package futures

import (
	"fmt"
	"sync"
	"time"
)

type Completer <-chan interface{}

type Future struct {
	triggered bool
	item      interface{}
	err       error
	lock      sync.Mutex
	wg        sync.WaitGroup
}

//
func (f *Future) GetResult() (interface{}, error) {
	f.lock.Lock()
	if f.triggered {
		f.lock.Unlock()
		return f.item, f.err
	}
	f.lock.Unlock()

	f.wg.Wait()
	return f.item, f.err
}

func (f *Future) HasResult() bool {
	f.lock.Lock()
	hasResult := f.triggered
	f.lock.Unlock()
	return hasResult
}

func (f *Future) setItem(item interface{}, err error) {
	f.lock.Lock()
	f.triggered = true
	f.item = item
	f.err = err
	f.lock.Unlock()
	f.wg.Done()
}

func listenForResult(f *Future, ch Completer, timeout time.Duration, wg *sync.WaitGroup) {
	wg.Done()
	t := time.NewTimer(timeout)
	select {
	case item := <-ch:
		f.setItem(item, nil)
		// we want to trigger GC of this timer as soon as it's no longger needed
		t.Stop()
	case <-t.C:
		f.setItem(nil, fmt.Errorf(`timeout after %f seconds`, timeout.Seconds()))
	}
}

func New(completer Completer, timeout time.Duration) *Future {
	f := &Future{}
	f.wg.Add(1)
	var wg sync.WaitGroup
	wg.Add(1)
	go listenForResult(f, completer, timeout, &wg)
	wg.Wait()
	return f
}
