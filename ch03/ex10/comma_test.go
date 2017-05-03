package main

import (
	"testing"
)

var f = comma

func TestComma1(t *testing.T) {
	result := f("1234")
	if result != "1,234" {
		t.Errorf("Expected value is 1,234. Actual %s", result)
	}
}

func TestComma2(t *testing.T) {
	result := f("")
	if result != "" {
		t.Errorf("Expected value is . Actual %s", result)
	}
}

func TestComma3(t *testing.T) {
	result := f("12")
	if result != "12" {
		t.Errorf("Expected value is 12. Actual %s", result)
	}
}

func TestComma4(t *testing.T) {
	result := f("123")
	if result != "123" {
		t.Errorf("Expected value is 123. Actual %s", result)
	}
}
