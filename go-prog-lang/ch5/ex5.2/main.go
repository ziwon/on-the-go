package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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

	seen := make(map[string]int)
	temp := map[int][]string{}
	var a []int

	countNode(nil, doc, seen)

	// sorting by its value
	//  http://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
	//  What the hell is going on under the hood?? Siranai...-_-;;
	for k, v := range seen {
		temp[v] = append(temp[v], k)
	}

	for k := range temp {
		a = append(a, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range temp[k] {
			// if tab size is 8 spaces?
			if len(s)+1 > 8 {
				fmt.Printf("%s\t%d\n", s, k)
			} else {
				fmt.Printf("%s\t\t%d\n", s, k)
			}
		}
	}
}

func countNode(stack []string, n *html.Node, seen map[string]int) []string {
	if n.Type == html.ElementNode {
		if seen[n.Data] += 1; seen[n.Data] == 0 {
			seen[n.Data] = 1
		}
		stack = append(stack, n.Data)
	}

	if n.FirstChild != nil {
		stack = countNode(stack, n.FirstChild, seen)
	}

	if n.NextSibling != nil {
		stack = countNode(stack, n.NextSibling, seen)
	}

	return stack
}

/*

╰─ go run main.go http://naver.com
span            388
a               236
div             135
li              80
strong          51
em              43
img             39
option          34
input           29
dd              28
p               28
ul              23
script          23
br              22
meta            15
h3              7
iframe          7
h2              7
dt              7
dl              7
button          5
label           5
hr              4
h4              4
link            4
area            4
fieldset        3
legend          3
optgroup        2
map             2
form            2
body            1
ol              1
h1              1
head            1
title           1
noscript        1
address         1
html            1
select          1
*/
