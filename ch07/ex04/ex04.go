package main

import (
	"io"

	"fmt"

	"golang.org/x/net/html"
)

func main() {
	node, err := Parse("<html><title>test</title></html>")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(node.FirstChild.Data)
}

//Parse return html, type of string, node.
func Parse(s string) (*html.Node, error) {
	return html.Parse(NewReader(s))
}

type reader struct {
	content []byte
	next    int
}

//Readは引数のバイトスライスにreaderが
func (r *reader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.next >= len(r.content) {
		return 0, io.EOF
	}

	nBytes := len(r.content) - r.next
	if nBytes > len(p) {
		nBytes = len(p)
	}

	copy(p, r.content[r.next:r.next+nBytes]) //読み込む分コピーする
	r.next += nBytes                         //読み込んだ分次に進める
	return nBytes, nil
}

//NewReaderはsを読み出すReaderを返す
func NewReader(s string) io.Reader {
	return &reader{[]byte(s), 0}
}
