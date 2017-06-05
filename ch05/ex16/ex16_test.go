package main

import "testing"

func TestEx16(t *testing.T) {
	actual := join("?", "a", "b", "c")
	expected := "a?b?c"
	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}

	actual = join("", "a", "b", "c")
	expected = "abc"
	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}
