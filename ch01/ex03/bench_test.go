package main

import (
	"os"
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	os.Args = []string{"test", "aaa", "bbb", "fasdfadsfaas", "fdafasdfad", "fdfasfa"}
	for i := 0; i < b.N; i++ {
		plusFunc()
	}
}

func BenchmarkJoin(b *testing.B) {
	os.Args = []string{"test", "aaa", "bbb", "fasdfadsfaas", "fdafasdfad", "fdfasfa"}
	for i := 0; i < b.N; i++ {
		joinFunc()
	}
}
