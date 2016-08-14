package reader

import (
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func TestReader(t *testing.T) {
	doc, err := html.Parse(NewReader("<p>hello</p>"))
	if err != nil {
		t.Error(err)
	}
	outline(nil, doc)
}
