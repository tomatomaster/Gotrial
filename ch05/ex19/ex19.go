package main

import "fmt"

func main() {
	defer magic()
	r := recover()
	fmt.Print(r)
}

func magic() {
	if r := recover(); r != nil {

	}

	panic("panic")
}
