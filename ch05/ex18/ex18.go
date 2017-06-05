package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()
	name, n, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %v ", name, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer closeFile(f, &err)
	return local, n, err
}

func closeFile(f *os.File, rError *error) {
	if err := f.Close(); err != nil {
		*rError = err
	}
}
