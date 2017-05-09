package main

import "testing"
import "os"

var f = wordfreq

func TestWordFreq(t *testing.T) {
	file, _ := os.Open("./res/test.txt")
	result := wordfreq(file)
	expected := make(map[string]int)
	expected["apple"] = 2
	expected["banana"] = 1
	expected["orange"] = 3
	if !equal(result, expected) {
		t.Errorf("%v %v", result, expected)
	}
}

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
