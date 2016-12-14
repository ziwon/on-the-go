package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

var depth int

func crawl(url string, depth int) []string {
	if depth <= 0 {
		return []string{}
	}

	fmt.Println("fetching:", url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	aFlag := flag.NewFlagSet("", flag.ExitOnError)
	aFlag.IntVar(&depth, "depth", 0, "a depth of links that will be fetched")
	lastIndex := len(os.Args) - 1
	aFlag.Parse(os.Args[lastIndex:])

	go func() {
		worklist <- os.Args[1:lastIndex]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, depth)
				depth -= 1
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			fmt.Println("> link", link)
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
