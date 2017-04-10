package main

import (
	"testing"

	"./popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(3)
	}
}

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountOriginal(3)
	}
}
