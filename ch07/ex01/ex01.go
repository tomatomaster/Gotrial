package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCount int
type ByteCounter int

func main() {
	var c WordCount
	fmt.Fprintf(&c, "hello, %s\n Today, I will be in here. \n This is", "Dolly")
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCount) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		*c += WordCount(1)
	}
	return len(p), nil
}
