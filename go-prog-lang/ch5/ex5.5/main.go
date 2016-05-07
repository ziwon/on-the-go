package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)

		if err == nil {
			fmt.Printf("words: %d, images: %d", words, images)
		}
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 0, 0, err
	}
	words, images = countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	var w, i int

	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		// How to count unique words? Well, I don't care.
		w += countWords(n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		i += 1
	}

	if n.FirstChild != nil {
		words, images := countWordsAndImages(n.FirstChild)
		w += words
		i += images
	}

	if n.NextSibling != nil {
		words, images := countWordsAndImages(n.NextSibling)
		w += words
		i += images
	}

	return w, i
}

func countWords(input string) (count int) {
	// pre-processing for two length whitespace, new line and tab
	input = strings.Replace(input, "  ", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\t", "", -1)

	if len(input) == 0 {
		return 0
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Println(input, count)
	return count
}
