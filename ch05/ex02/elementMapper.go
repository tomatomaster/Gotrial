package main

import (
	"fmt"
	"os"

	"github.com/tomatomaster/gotorial/gotorial"
	"golang.org/x/net/html"
)

func dummymain() {
	doc, err := html.Parse(os.Stdin)
	gotorial.OSExitIfError(err)
	m := make(map[string]int)
	m = mapper(m, doc)

	for k, v := range m {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

func mapper(m map[string]int, n *html.Node) (r map[string]int) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = mapper(m, c)
	}
	return m
}
