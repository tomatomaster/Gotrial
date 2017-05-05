package main

import (
	"fmt"
)

func main() {
	s := []string{"s", "s", "s", "o", "s", "a", "b", "v", "v", "x"}
	fmt.Println(s)
	s = eliminate(s)
	fmt.Println(s)
}

func eliminate(s []string) []string {
	for i := 0; i < len(s); {
		if i+1 < len(s) && s[i] == s[i+1] {
			s = remove(s, i+1)
		} else {
			i++
		}
	}
	return s
}

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
