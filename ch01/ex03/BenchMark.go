package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	bench(plusFunc)
	bench(joinFunc)
}

func bench(fn func()) {
	tS := time.Now()
	fn()
	tE := time.Now()
	fmt.Printf("TIME: %d \n", tE.Nanosecond()-tS.Nanosecond())
}

func plusFunc() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func joinFunc() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
