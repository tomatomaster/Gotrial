package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	target := []byte("今日は   良い     天気    ですね。")
	compressed := comp(target)
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
				count += size
			} else {
				i += size
			}
		} else {
			i += size
		}
	}
	return b[:len(b)-count]
}

func comp(b []byte) []byte {
	for i := 0; i < len(b); {
		r1, size1 := utf8.DecodeRune(b[i:])
		r2, size2 := utf8.DecodeRune(b[i+size1:])
		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			b[i] = ' '
			copy(b[i:], b[(i+size1+size2):])
		}
		i += size1 + size2
	}
	return b
}

func test(s string) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) //Decodeしたい文字だけを渡してもマルチバイトなので、文字の終端を判断できない
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
