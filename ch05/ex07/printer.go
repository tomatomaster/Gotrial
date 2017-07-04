package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, _ := http.Get(url)
	node, _ := html.Parse(resp.Body)
	forEachNode(node, startElement, endElement)
}

var depth int

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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attrs []string
		for _, attr := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s=\"%s\"", attr.Key, attr.Val))
		}
		if n.FirstChild == nil {
			if attrs != nil {
				fmt.Printf("%*s<%s %s/>", depth*2, "", n.Data, strings.Join(attrs, " "))
			} else {
				fmt.Printf("%*s<%s/>", depth*2, "", n.Data)
			}
		} else {
			if attrs != nil {
				fmt.Printf("%*s<%s %s>", depth*2, "", n.Data, strings.Join(attrs, " "))
			} else {
				fmt.Printf("%*s<%s>", depth*2, "", n.Data)
			}
			depth++
		}
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Printf("<!--%s-->", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.FirstChild == nil {
		return
	}
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>", depth*2, "", n.Data)
	}
}
