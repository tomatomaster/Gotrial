package main

import (
	"testing"
)

var f = charcount

func TestCharcout1(t *testing.T) {
	result := charcount('O')
	expected := []string{"Graphic", "Print", "Letter", "Upper"}
	if !equal(result, expected) {
		t.Errorf("%v %v", result, expected)
	}
}

func TestCharcout2(t *testing.T) {
	result := charcount('💏')
	expected := []string{"Graphic", "Print", "Symbol"}
	if !equal(result, expected) {
		t.Errorf("%v %v", result, expected)
	}
}

func TestCharcout3(t *testing.T) {
	result := charcount('1')
	expected := []string{"Graphic", "Print", "Digit", "Number"}
	if !equal(result, expected) {
		t.Errorf("%v %v", result, expected)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	emap := make(map[string]bool)
	for _, ea := range a {
		emap[ea] = true
	}
	for _, eb := range b {
		if emap[eb] == false {
			return false
		}
	}
	return true
}
