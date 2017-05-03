package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "test"
	s2 := "stet"
	fmt.Printf("Compare %s to %s\n", s1, s2)
	r := isAnagram(s1, s2)
	fmt.Printf("insAnagram %v\n", r)
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var map1 = make(map[rune]int)
	for i := 0; i < utf8.RuneCountInString(s1); i++ {
		e, _ := utf8.DecodeRuneInString(s1[i:])
		map1[e]++
	}

	for i := 0; i < utf8.RuneCountInString(s2); i++ {
		e, _ := utf8.DecodeRuneInString(s2[i:])
		if map1[e] == 0 {
			return false
		}
		map1[e]--
	}

	return true
}
