package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo()
}

func echo() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
