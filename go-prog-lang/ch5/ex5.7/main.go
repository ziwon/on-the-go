package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		depth++
	} else if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, extractAttributes(n))
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.TextNode {
		depth--
	} else if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func extractAttributes(n *html.Node) string {
	var i int
	var len = len(n.Attr)

	buffer := new(bytes.Buffer)
	for _, a := range n.Attr {
		buffer.WriteString(a.Key)
		buffer.WriteString("=")
		buffer.WriteString("\"")
		buffer.WriteString(a.Val)
		buffer.WriteString("\"")
		if i != len-1 {
			buffer.WriteString(" ")
			i++
		}
	}
	return buffer.String()
}
