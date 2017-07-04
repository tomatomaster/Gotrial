// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster() //Clientからのイベントを受信するゴルーチン

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

//Client
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- fmt.Sprintf("You are %s", who)
	messages <- fmt.Sprintf("%s has arrived", who)
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("%s: %s", who, input.Text())
	}

	leaving <- client{ch, who}
	messages <- fmt.Sprintf("%s has left", who)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

//Broadcast
type client struct {
	receivedMsg chan<- string
	userName    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.receivedMsg <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			for c := range clients {
				cli.receivedMsg <- fmt.Sprintf("%s login room", c.userName)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.receivedMsg)
		}
	}
}
