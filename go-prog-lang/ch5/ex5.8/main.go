package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	outline(os.Args[1], os.Args[2])
}

func outline(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	el := ElementByID(doc, id)
	if el == nil {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found")
	}

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) {
			return true
		}
	}

	if post != nil {
		if post(n) {
			return true
		}
	}

	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node

	forEachNode(doc,
		func(n *html.Node) bool {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					found = n
					return true
				}
			}
			return false
		},
		func(n *html.Node) bool {
			return false
		})

	return found
}
