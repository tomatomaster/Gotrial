package ex17

import (
	"golang.org/x/net/html"
)

func ElementByTagName(n *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	forEachNode(n, func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, data := range n.Data {
				for _, n := range name {
					if data == n {
						nodes = append(nodes, n)
					}
				}
			}
		}
	}, nil)
	return nodes
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
