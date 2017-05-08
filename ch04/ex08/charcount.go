package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	for _, r := range "Goがチャームポイントです🤗💗💗" {
		result := charcount(r)
		for _, category := range result {
			counts[category]++
		}
		fmt.Printf("%s %v\n", string(r), result)
	}
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func charcount(r rune) []string {
	var category []string

	category = loop(category, r, unicode.IsControl)
	category = loop(category, r, unicode.IsDigit)
	category = loop(category, r, unicode.IsGraphic)
	category = loop(category, r, unicode.IsLetter)
	category = loop(category, r, unicode.IsLower)
	category = loop(category, r, unicode.IsMark)
	category = loop(category, r, unicode.IsNumber)
	category = loop(category, r, unicode.IsPrint)
	category = loop(category, r, unicode.IsPunct)
	category = loop(category, r, unicode.IsSpace)
	category = loop(category, r, unicode.IsSymbol)
	category = loop(category, r, unicode.IsTitle)
	category = loop(category, r, unicode.IsUpper)

	return category
}

func loop(category []string, r rune, isMethod func(rune) bool) []string {
	if isMethod(r) {
		return append(category, getFunctionName(isMethod))
	}
	return category
}

func getFunctionName(f interface{}) string {
	fv := reflect.ValueOf(f)
	hasIs := runtime.FuncForPC(fv.Pointer()).Name()
	return strings.TrimPrefix(hasIs, "unicode.Is")
}
