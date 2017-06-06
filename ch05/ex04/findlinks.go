package main

import "golang.org/x/net/html"
import "os"
import "github.com/tomatomaster/gotorial/gotorial"
import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get(os.Args[1])
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	gotorial.OSExitIfError(err)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

var dataSet map[string]string = map[string]string{"a": "href", "img": "src", "script": "src"}

func visit(links []string, n *html.Node) []string {
	for k, v := range dataSet {
		links = appendLink(links, n, k, v)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func appendLink(links []string, n *html.Node, data, attr string) []string {
	if n.Type == html.ElementNode && n.Data == data {
		for _, a := range n.Attr {
			if a.Key == attr {
				links = append(links, a.Val)
			}
		}
	}
	return links
}

func vi(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if next := n.FirstChild; next != nil {
		links = vi(links, next)
	}
	if next := n.NextSibling; next != nil {
		links = vi(links, next)
	}
	return links
}

func v(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = v(links, n.FirstChild)
	links = v(links, n.NextSibling)

	return links
}
