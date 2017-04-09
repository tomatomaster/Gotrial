package main

import (
	"testing"

	"./pcount"
)

var pop = popCount

func BenchmarkOriginalPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pcount.PopCount(1023)
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pop(1023)
	}
}

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
