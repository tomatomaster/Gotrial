package main

import (
	"fmt"
	"os"

	"log"

	"regexp"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open(`./sample.html`)
	if err != nil {
		log.Fatal(err)
	}
	node, _ := html.Parse(f)
	showDocumentNode(node)
}

func showDocumentNode(node *html.Node) {
	r := regexp.MustCompile(`script|style`)
	if node.Type == html.ElementNode {
		if !r.MatchString(node.Data) {
			if node.FirstChild.Type == html.TextNode {
				fmt.Println(node.FirstChild.Data)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		showDocumentNode(c)
	}
}
