// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.Tick(1 * time.Second)
	connection := make(chan net.Conn)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connection <- conn
	}()

	countDown := 10
	for {
		select {
		case c := <-connection:
			handleConn(c)
			countDown = 10
		case <-ticker:
			countDown--
			if countDown == 0 {
				fmt.Println("Disconnected")
				return
			}
		}
	}

}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
