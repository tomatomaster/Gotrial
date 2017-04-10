package main

import (
	"fmt"

	"./popcount"
)

func main() {
	fmt.Printf("Set Bit: %d", popcount.PopCount(2047))
}
