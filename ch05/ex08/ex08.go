package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	node, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	findNode := ElementByID(node, "class")
	fmt.Printf("%v\n", findNode.Data)
}

func ElementByID(n *html.Node, id string) *html.Node {
	var node *html.Node
	forEachNode(n, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == id {
					node = n
					return false
				}
			}
		}
		return true
	}, nil)
	return node
}

var ok bool = true

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if ok == false {
		return
	}

	if pre != nil {
		ok = pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		ok = post(n)
	}
}
