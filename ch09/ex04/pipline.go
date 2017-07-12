// Copyright © 2017 Ryutarou Ono.

package main

import "fmt"

var done = make(chan struct{})

func main() {
	pipeline(10)
	<-done
}

func pipeline(num int) {
	next := make(chan int)
	for i := 0; i < num; i++ {
		from := make(chan int)
		from = next
		go func(from chan int) {
			if i == 0 {
				next <- 0
			} else {
				f := <-from
				f++
				fmt.Printf("from: %d\n", f)
				next <- f
			}
		}(from)
	}
	done <- struct{}{}
}
