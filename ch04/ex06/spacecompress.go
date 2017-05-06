package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	target := []byte("今日は   良い     天気    ですね。")
	compressed := compress(target)
	var stringval = string(compressed)
	fmt.Printf("%s \n", stringval)
}

func compress(b []byte) []byte {
	var count int
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			next := i + size
			nRune, _ := utf8.DecodeRune(b[next:])
			if unicode.IsSpace(nRune) {
				copy(b[i:], b[next:])
				count++
			} else {
				i += size
			}
		} else {
			i += size
		}
	}
	return b[:len(b)-count]
}

func test(s string) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) //Decodeしたい文字だけを渡してもマルチバイトなので、文字の終端を判断できない
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
