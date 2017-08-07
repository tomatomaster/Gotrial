// Copyright © 2017 Ryutarou Ono.

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	tcp, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("not OK")
	}
	go func() {
		_, err = io.Copy(os.Stdout, tcp)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(tcp, os.Stdin)
	os.Stdin.Close()
	tcp.CloseWrite() //tcp.Readは生きてるので、
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
