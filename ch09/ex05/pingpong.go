// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"time"
)

var pinA = make(chan string)
var pinB = make(chan string)

func main() {

	go func() {
		var count int
		for {
			mmseg := <-pinB
			count++
			fmt.Printf("Receive mmesg %s : %d\n", mmseg, count)
			pinA <- "ping !"
		}
	}()

	go func() {
		for {
			mmseg := <-pinA
			fmt.Printf("Receive mmesg %s\n", mmseg)
			pinB <- "pong !"
		}
	}()

	pinB <- "First"
	time := time.After(1 * time.Second)
	<-time
}
