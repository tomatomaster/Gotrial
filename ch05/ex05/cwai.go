package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	fmt.Printf("url %s\n", url)
	words, images, _ := CountWordsAndImages(url)
	fmt.Printf("word %d image %d \n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	if doc.Type == html.ElementNode && doc.Data == "img" {
		images++
	} else if doc.Type == html.TextNode {
		words += len(strings.Split(doc.Data, " "))
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
