// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

type palind string

func (s palind) Len() int           { return utf8.RuneCountInString(string(s)) }
func (s palind) Swap(i, j int)      {}
func (s palind) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	var p1 palind
	p1 = "test"
	var p2 palind
	p2 = "Funky yknuF"
	fmt.Printf("IsPalindrome(%s) = %v\n", p1, IsPalindrome(p1))

	fmt.Printf("IsPalindrome(%s) = %v\n", p2, IsPalindrome(p2))
}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()-1; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}
