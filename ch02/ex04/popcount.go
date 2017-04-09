package main

import (
	"fmt"
)

func main() {
	count := popCount(1023)
	fmt.Print(count)
}

func popCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}
