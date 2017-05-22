package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(comma("345234413143424267"))
}

func comma(s string) string {
	if len(s) == 3 {
		return s
	}
	b := []byte(s)
	var buf bytes.Buffer
	count := 0
	mod := len(b) % 3
	for i := 0; i < mod; i++ {
		buf.WriteByte(b[count])
		count++
	}
	//fmt.Printf("%d %d \n", count, mod) // 1 1 両方とも同じ値を指すにもかかわらず、countを使用すると、1,234でforが止まる。なぜ？
	for i := 0; i < len(b)-mod; i++ { // for i := 0; i < len(b)-count; i++
		if i%3 == 0 && len(b) != 3 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[count])
		count++
	}
	return buf.String()
}
