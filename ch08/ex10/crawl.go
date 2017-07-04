// Copyright © 2017 Ryutarou Ono.

package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"golang.org/x/net/html"
)

var done = make(chan struct{})

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func() { worklist <- os.Args[1:] }()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Printf("Interrupted Cancel Signel\n")
		close(done)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if cancelled() {
					return
				}
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := cancellableExtract(url)
	<-tokens
	if err != nil {
		log.Fatal(err)
	}
	return list
}

/**
Cancell
*/

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// Cf cancel http://qiita.com/atijust/items/63676309c7b3d5df5948
func cancellableExtract(url string) (hrefs []string, err error) {
	newReq, err := http.NewRequest("GET", url, nil)
	cancel := make(chan struct{})
	newReq.Cancel = cancel
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Received Cancel Signal. Then Close Request")
				close(cancel)
			}
		}
	}()
	resp, err := http.DefaultClient.Do(newReq)
	defer resp.Body.Close()
	hrefs = extractLinks(resp)

	doc, _ := html.Parse(resp.Body)
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				hrefs = append(hrefs, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return hrefs, err
}

func extractLinks(resp *http.Response) (links []string) {
	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					link, err := resp.Request.URL.Parse(attr.Val)
					if err != nil {
						log.Fatal(err)
					}
					links = append(links, link.String())
				} else {
					continue
				}
			}
		}
	}
	return links
}

// Copied from gopl.io/ch5/outline2.
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
