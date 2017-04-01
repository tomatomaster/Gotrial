package main

import (
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusFunc()
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinFunc()
	}
}
