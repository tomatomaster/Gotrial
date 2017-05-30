package main

import (
	"fmt"
	"regexp"
)

func main() {
	result := expand("Test food taste good,", func(s string) string {
		return "hogehoge"
	})
	fmt.Println(result)
}

func expand(s string, f func(string) string) string {
	seed := f("foo")
	rep := regexp.MustCompile(`foo`)
	return rep.ReplaceAllString(s, seed)
}
