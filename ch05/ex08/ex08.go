package main

import (
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, _ := http.Get(url)
	html.Parse(resp.Body)
}

func ElementByID(n *html.Node, id string) *html.Node {
	forEachNode(n, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return false
				}
			}
		}
		return true
	}, nil)
	return n
}

func find(id string, n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
	}
	return true
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
