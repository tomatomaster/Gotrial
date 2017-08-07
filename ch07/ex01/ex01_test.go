package main

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	var counter WordCount
	fmt.Fprintf(&counter, "Test 4 Word %s", "!")
	if counter != WordCount(4) {
		t.Fatalf("Expected val 4, but %v", counter)
	}
}
