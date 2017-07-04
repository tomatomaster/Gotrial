// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"gopl.io/ch5/links"
)

type listInf struct {
	link  string
	depth int
}

func main() {
	var depth int
	flag.IntVar(&depth, "depth", -1, "Indicates search depth")
	flag.Parse()

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			var linkDepth int
			if depth == -1 {
				for link := range unseenLinks {
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
			for link := range unseenLinks {
				if linkDepth < depth {
					linkDepth++
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
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

//!+crawl from gopl.io
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
