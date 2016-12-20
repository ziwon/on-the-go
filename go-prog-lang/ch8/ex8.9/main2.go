package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]

	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fn := make(chan func() (string, int64))
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fn)
	}

	go func() {
		n.Wait()
		close(fn)
	}()

	tick := time.Tick(500 * time.Millisecond)
	nbytes := make(map[string]int64, len(roots))

loop:

	for {
		select {
		case <-done:
			for range fn {
				// Do nothing
			}
			return
		case call, ok := <-fn:
			if !ok {
				break loop
			}
			root, size := call()
			nbytes[root] += size
		case <-tick:
			printDiskUsage2(nbytes)
		}
	}
	printDiskUsage2(nbytes)
}

func printDiskUsage2(nbytes map[string]int64) {
	for k, v := range nbytes {
		fmt.Printf("%s  %.1f GB\n", k, float64(v)/1e9)
	}
}

func walkDir(root, dir string, n *sync.WaitGroup, fn chan<- func() (string, int64)) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, infos)
		} else {
			fn <- (func() (string, int64) { return root, entry.Size() })
		}
	}
}

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
