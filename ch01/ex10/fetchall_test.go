package main

import (
	"os"
	"testing"
)

var f = writeFileTo

func TestWriteFileTo(t *testing.T) {
	file, error := os.Open("./test/testSet")
	if error != nil {
		t.Error("Can't open file")
	}
	defer file.Close()
	_, err := f("./test/test", file)
	if err != nil {
		t.Error("Can't open file")
	}
}
