// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const protocol = "tcp4"

func main() {
	listener, err := net.Listen(protocol, "localhost:8000")
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

const capability = 10
const connectionTimeOut = 5 * time.Minute

//Client
func handleConn(conn net.Conn) {
	ch := make(chan string, capability)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- fmt.Sprintf("You are %s", who)
	messages <- fmt.Sprintf("%s has arrived", who)
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)

	done := make(chan struct{})
	//If Client sleeps for five minutes, server closes connection.
	c := time.Tick(connectionTimeOut)
	go func(conn net.Conn) {
		for {
			select {
			case <-c:
				messages <- fmt.Sprint("Close Connection")
				conn.Close()
			case <-done:
				c = time.Tick(connectionTimeOut)
			}
		}
	}(conn)

	/*
		timer := time.AfterFunc(3000*time.Millisecond, func() {
			conn.Close()
		})
	*/

	for input.Scan() {
		//timer.Stop()
		messages <- fmt.Sprintf("%s: %s", protocol, input.Text())
		/*		timer = time.AfterFunc(10*time.Second, func() {
				conn.Close()
			})*/
		done <- struct{}{}
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
