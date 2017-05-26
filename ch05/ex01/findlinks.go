package main

import "golang.org/x/net/html"
import "os"
import "github.com/tomatomaster/gotorial/gotorial"
import "fmt"

func main() {
	doc, err := html.Parse(os.Stdin)
	gotorial.OSExitIfError(err)
	for _, link := range visit(nil, doc) {
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
}


