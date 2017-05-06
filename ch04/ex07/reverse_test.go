package main

import (
	"testing"
)

var f = reverse

func TestReverse(t *testing.T) {
	sample := []byte("今日はいい日だな")
	expected := "なだ日いいは日今"
	result := f(sample)
	if expected != string(result) {
		t.Errorf("Expect: %s Actual: %s", expected, result)
	}
}

func TestReverse2(t *testing.T) {
	sample := []byte("Hello")
	expected := "olleH"
	result := f(sample)
	if expected != string(result) {
		t.Errorf("Expect: %s Actual: %s", expected, result)
	}
}
