package main

import (
	"fmt"
)

func main() {
	months := [...]string{1:"January", 12: "December"}
	months_slice := months[1:13]

	for i, v := range months {
		fmt.Printf("%d %s\n",i, v)
	}
	fmt.Println()
	for i, v := range months_slice {
		fmt.Printf("%d %s\n",i, v)
	}

	fmt.Printf("%v %v",cap([]int{}), len([]int{}))
	slice := make([]string, 5, 10)
	fmt.Print(len(slice))

	var runes []rune
	for _, r:= range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}