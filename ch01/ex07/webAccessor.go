package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		PrintRespBody(url)
	}
}

//PrintRespBody response body accessing url.
func PrintRespBody(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "", err)
		return false
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return false
	}
	return true
}
