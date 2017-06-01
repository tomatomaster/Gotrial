package main

import (
	"log"

	"os"

	"strings"

	"net/http"

	"io/ioutil"

	"fmt"

	"./links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item, domain string) []string, worklist []string) {
	rootWork := worklist[0]
	domain := strings.TrimSuffix(strings.SplitAfter(rootWork, "/")[2], "/")
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !(seen[item]) {
				seen[item] = true
				worklist = append(worklist, f(item, domain)...)
			}
		}
	}
}

func crawl(url, domain string) []string {
	fmt.Println(url)
	urlDomain := strings.TrimSuffix(strings.SplitAfter(url, "/")[2], "/")
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, value := range list {
		if urlDomain == domain {
			writePage(value)
		}
	}
	return list
}

func writePage(url string) {
	file, err := os.Create("./t.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(body)
}

func isSameDomain(url, domain string) bool {
	urlDomain := strings.TrimSuffix(strings.SplitAfter(url, "/")[2], "/")
	if domain == urlDomain {
		return true
	}
	return false
}
