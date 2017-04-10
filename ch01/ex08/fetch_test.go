package main

import (
	"testing"
)

var testFunc = ifNotHasPrefixThenAppend

func TestIfNotHasPrefixThenAppend(t *testing.T) {
	result := testFunc("prefix:", "test")
	if result != "prefix:test" {
		t.Errorf("Faile: expected prefix:test, but actual %s", result)
	}
}
