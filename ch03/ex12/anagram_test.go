package main

import (
	"testing"
)

var f = isAnagram

func TestIsAnagram1(t *testing.T) {
	s1 := "a"
	s2 := "b"
	r := f(s1, s2)
	if r == true {
		t.Errorf("Expected value is false but actual %v", r)
	}
}

func TestIsAnagram2(t *testing.T) {
	s1 := "test"
	s2 := "tset"
	r := f(s1, s2)
	if r == false {
		t.Errorf("Expected value is true but actual %v", r)
	}
}

func TestIsAnagram3(t *testing.T) {
	s1 := "ttest"
	s2 := "test"
	r := f(s1, s2)
	if r == true {
		t.Errorf("Expected value is false but actual %v", r)
	}
}
