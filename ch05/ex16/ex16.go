package main

import (
	"fmt"
	"strings"
)

func main() {
	r := join("+", "a", "b", "C")
	fmt.Print(r)
}

func join(sep string, s ...string) string {
	return strings.Join(s, sep)
}
