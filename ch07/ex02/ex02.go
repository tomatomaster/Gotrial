package main

import (
	"fmt"
	"io"
)

func main() {
	writer, counter := CountingWriter(bytes)
	fmt.Println(writer)
	fmt.Println(counter)
}

type OriWriter struct {
	w io.Writer
	b int64
}

func (ow *OriWriter) Write(p []byte) (int, error) {
	ow.b += int64(len(p))
	return ow.w.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var owriter OriWriter
	owriter.w = w
	return owriter.w, &owriter.b
}
