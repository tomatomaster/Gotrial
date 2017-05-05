package main

import (
	"fmt"
)

func main() {
	val := [32]int{1, 2, 3, 4}
	rev := reverse(&val)
	fmt.Println(rev)
}

func reverse(ptr *[32]int) *[32]int {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
	return ptr
}
