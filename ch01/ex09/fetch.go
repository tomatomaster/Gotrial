package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = ifNotHasPrefixThenAppend("http://", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "", err)
			os.Exit(1)
		}
		printStatusCode(os.Stdout, resp)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func printStatusCode(w io.Writer, resp *http.Response) {
	fmt.Fprintln(w, "StatusCode:"+resp.Status)
}

func ifNotHasPrefixThenAppend(prefix string, dest string) string {
	if !strings.HasPrefix(dest, "http://") {
		dest = "http://" + dest
	}
	return dest
}
