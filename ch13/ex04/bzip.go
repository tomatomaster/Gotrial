// Copyright © 2017 Ryutarou Ono.
// ????
package main

import (
	"io"
	"os"
	"os/exec"
	"sync"
)

func main() {
	zip, _ := os.OpenFile("comp.bz2", os.O_WRONLY|os.O_CREATE, 0666)
	f, _ := os.Open("test.txt")
	writer := NewWriter(zip)
	buf := make([]byte, 1024)
	f.Read(buf)
	writer.Write(buf)
	writer.Close()
}

type writer struct {
	mu sync.Mutex
	w  io.WriteCloser
}

func NewWriter(out io.Writer) io.WriteCloser {
	var writer writer
	cmd := exec.Command("/usr/bin/bzip2")
	pReader, pWriter := io.Pipe()
	cmd.Stdin = pReader
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		io.Copy(out, pReader)
		w.Done()
	}()
	cmd.Start()
	writer.w = pWriter
	return writer.w
}

func (w *writer) Write(data []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	var total int
	for len(data) > 0 {
		n, _ := w.w.Write(data)
		total += n
		data = data[total:]
	}
	return total, nil
}

func (w *writer) Close() {
	w.mu.Lock()
	w.mu.Unlock()
	w.w.Close()
}
