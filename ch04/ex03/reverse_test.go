package main

import (
	"testing"
)

var f = reverse

func TestReverse(t *testing.T) {
	val := [32]int{1, 2, 3, 4, 5, 6}
	expected := [32]int{26: 6, 27: 5, 28: 4, 29: 3, 30: 2, 31: 1}
	result := reverse(&val)
	if val != expected {
		t.Errorf("\nResult  :%d\nExpected:%d", *result, expected)
	}
}
