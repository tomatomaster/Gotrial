package main

import "golang.org/x/net/html"
import "os"
import "github.com/tomatomaster/gotorial/gotorial"
import "fmt"

func main() {
	doc, err := html.Parse(os.Stdin)
	gotorial.OSExitIfError(err)
	for _, link := range vi(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(nil, c)
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
