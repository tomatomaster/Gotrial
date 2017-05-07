package main

import (
	"fmt"
)

var a = 0

func init() {
	a = 1
}

func init() {
	fmt.Print("yatta")
}

func init() {
	fmt.Print("test")
}

func main() {
	fmt.Print(a)
}
