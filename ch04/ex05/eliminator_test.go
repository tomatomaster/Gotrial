package main

import (
	"testing"
)

var f = eliminate

func Test1(t *testing.T) {
	sample := []string{"z", "z", "h", "e", "l", "l", "o", "a"}
	result := f(sample)

	var temp string
	for _, e := range result {
		if temp != "" && temp == e {
			t.Error("Find duplicate value")
		}
		temp = e
	}
}

func Test2(t *testing.T) {
	sample := []string{"z", "z", "h", "e", "l", "z", "z", "l", "o", "a"}
	result := f(sample)

	var temp string
	for _, e := range result {
		if temp != "" && temp == e {
			t.Error("Find duplicate value")
		}
		temp = e
	}
}
func Test3(t *testing.T) {
	sample := []string{}
	result := f(sample)

	var temp string
	for _, e := range result {
		if temp != "" && temp == e {
			t.Error("Find duplicate value")
		}
		temp = e
	}
}
