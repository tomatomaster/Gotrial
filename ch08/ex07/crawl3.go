// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"net/http"
	"path/filepath"

	"strings"

	"bufio"

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
					writeToLocal(link)
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
			for link := range unseenLinks {
				if linkDepth < depth {
					linkDepth++
					writeToLocal(link)
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

var rootDir = getCurrentDir()

func writeToLocal(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//TopLevel Domain
	if len(strings.Split(url, "/")) == 3 {
		// https://google.com -> google.com
		hostName := getHostName(url)
		file := createIndexPage(hostName)
		bufio.NewReader(resp.Body).WriteTo(file)
	} else {
		subName := getSubName(url)
		err = os.MkdirAll(subName, 0777)
		if err != nil {
			log.Fatal(err)
		}
		rootPath := filepath.Dir("./")
		file, err := os.Create(filepath.Join(rootPath, subName, "index.html"))
		if err != nil {
			log.Fatal(err)
		}
	}


	else if !strings.Contains(url, "/") {
		//rootDir
		err = os.MkdirAll(url, 0777)

		file, err := os.OpenFile(filepath.Join(url, "root"), os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		bufio.NewReader(resp.Body).WriteTo(file)
	} else {
		err = os.MkdirAll(url, 0600)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.OpenFile(url, os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		bufio.NewReader(resp.Body).WriteTo(file)
	}

}

//TODO https://google.com returns google.com
func getHostName(url string) string {
	return strings.Split(url, "/")[2]
}

//TODO https://google.com/sub/subsub returns sub/subsub
func getSubName(url string) string {
	url = strings.TrimSuffix(url, "http://")
	url = strings.TrimSuffix(url, "https://")
	return url
}

func createIndexPage(host string) *os.File {
	rootPath := filepath.Join("./", host)
	err := os.MkdirAll(rootPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(filepath.Join(rootPath, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
