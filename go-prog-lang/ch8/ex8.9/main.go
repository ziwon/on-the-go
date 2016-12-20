package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/**
go run main.go /home/luno/Workspace /home/luno/Downloads
/home/luno/Downloads  17.3 GB
/home/luno/Workspace  1.2 GB
/home/luno/Workspace  2.5 GB
/home/luno/Downloads  17.3 GB
/home/luno/Downloads  17.3 GB
/home/luno/Workspace  3.1 GB
/home/luno/Downloads  17.3 GB
/home/luno/Workspace  3.9 GB
/home/luno/Downloads  17.3 GB
/home/luno/Workspace  4.0 GB
*/

type fileInfo struct {
	root string
	size int64
}

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

	fileInfos := make(chan fileInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileInfos)
	}

	go func() {
		n.Wait()
		close(fileInfos)
	}()

	tick := time.Tick(500 * time.Millisecond)
	nbytes := make(map[string]int64, len(roots))

loop:

	for {
		select {
		case <-done:
			for range fileInfos {
				// Do nothing
			}
			return
		case infos, ok := <-fileInfos:
			if !ok {
				break loop
			}
			nbytes[infos.root] += infos.size
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

func walkDir(root, dir string, n *sync.WaitGroup, infos chan<- fileInfo) {
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
			infos <- fileInfo{root, entry.Size()}
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
