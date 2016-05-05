package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := []string{}
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		s = append(s, string(b))
	}

	reader := strings.NewReader(strings.Join(s, ""))
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//for _, texts := range visit(nil, doc) {
	//fmt.Println(texts)
	//}

	visit(doc)

}

func visit(n *html.Node) {
	if isTextNode(n) {
		if n.Data != "span" {
			s := strings.TrimSpace(n.Data)
			s = strings.Trim(s, "\n")
			if len(s) != 0 {
				fmt.Printf("%s\n", s)
			}
		} else {
			visit(n.FirstChild)
		}
	}

	if n.FirstChild != nil {
		if !isScript(n.FirstChild) && !isStyle(n.FirstChild) {
			visit(n.FirstChild)
		}
	}

	if n.NextSibling != nil {
		if !isScript(n.NextSibling) && !isStyle(n.NextSibling) {
			visit(n.NextSibling)
		}
	}
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

func isScript(n *html.Node) bool {
	return n.Data == "script"
}

func isStyle(n *html.Node) bool {
	return n.Data == "style"
}
