package main

import (
	"bufio"
	"fmt"
)

type Count string
type ByteCounter int

func main() {
	var c ByteCounter
	fmt.Fprintf(&c, "hello, %s", "Dolly")
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *Count) Write(p []byte) (int, error) {

	bufio.ScanWords()
}
