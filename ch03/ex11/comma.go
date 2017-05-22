package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma2("+12313133.12313"))
	fmt.Println(comma2("-1234.5"))
}

func comma2(s string) string {
	var number string
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		number = s[1:]
	} else {
		number = s
	}
	if strings.Contains(number, ".") {
		v := strings.Split(number, ".")
		return comma(v[0]) + "." + v[1]
	}
	buf.WriteString(comma(number))
	return buf.String()
}

func comma(s string) string {
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
		if i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[count])
		count++
	}
	return buf.String()
}
