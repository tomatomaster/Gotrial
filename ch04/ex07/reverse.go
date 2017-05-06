package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	in := []byte("いい日だなぁ")
	out := reverse(in)
	fmt.Printf("%s\n", string(out))

	in = []byte("今日はtestをdoする")
	out = reverse(in)
	fmt.Printf("%s\n", string(out))
}

func reverse(s []byte) []byte {
	temp := make([]byte, len(s))
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		copy(temp[len(s)-i-size:len(s)-i], s[i:i+size])
		i += size
	}
	copy(s, temp)
	return s
}
