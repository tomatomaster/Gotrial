package main

import (
	"fmt"
)

func main() {
	t := []int{1, 2, 3, 4, 5}
	rotate(t)
	fmt.Print(t)
}

func rotate(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[len(s)-i-1]
	}
}
