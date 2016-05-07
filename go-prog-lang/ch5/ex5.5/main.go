package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
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
	if n.Type == html.TextNode {
		w += 1
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
