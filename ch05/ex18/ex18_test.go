package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestEx18(t *testing.T) {
	err := errors.New("Dummy")
	f, _ := os.Create("file")
	closeFile(f, &err)
	if fmt.Sprint(err) != "Dummy" {
		t.Fatalf("Expected Dummy actual %v", err)
	}
}
