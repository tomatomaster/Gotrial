package main

import (
	"testing"
)

var fP2K = p2K

func TestP2K(t *testing.T) {
	result, _ := fP2K(1)
	expected := 0.453592
	if result != expected {
		t.Errorf("Result: %v Expected %v", result, expected)
	}
}

func TestP2KNegativeValue(t *testing.T) {
	result, err := fP2K(-1)
	expected := 0.453592
	if err == nil {
		t.Errorf("fP2K not support negative value")
	}
	if result != 0 {
		t.Errorf("Result: %v Expected %v", result, expected)
	}
}
