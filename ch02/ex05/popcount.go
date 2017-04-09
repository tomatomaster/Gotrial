package main

import (
	"fmt"
)

func main() {
	count := popCount(1023)
	fmt.Println(count)
}

func popCount(x uint64) int {
	count := 0
	for ; x != 0; x = x & (x - 1) {
		count++
	}
	return count
}
