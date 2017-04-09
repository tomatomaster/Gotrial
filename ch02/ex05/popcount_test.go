package main

import "testing"

var pop = popCount

func TestPopcount(t *testing.T) {
	if pop(0) != 0 {
		t.Errorf("UnExpected")
	}

	if pop(63) != 6 {
		t.Errorf("UnExpected")
	}

	if pop(1023) != 10 {
		t.Errorf("UnExpected")
	}
}
