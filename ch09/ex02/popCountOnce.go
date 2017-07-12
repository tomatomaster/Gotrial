// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"./popcount"
)

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Printf("Popcount %d\n", popcount.PopCount(10))
		done <- struct{}{}
	}()
	go func() {
		fmt.Printf("Popcount %d\n", popcount.PopCount(20))
		done <- struct{}{}
	}()
	go func() {
		fmt.Printf("Popcount %d\n", popcount.PopCount(100))
		done <- struct{}{}
	}()
	<-done
	<-done
	<-done
}
