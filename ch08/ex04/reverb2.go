// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}
	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
}

func echo(c net.Conn, text string, delay time.Duration, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
