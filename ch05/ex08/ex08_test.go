package main

import (
	"os"
	"testing"

	"github.com/tomatomaster/gotorial/gotorial"
	"golang.org/x/net/html"
)

func TestEx08(t *testing.T) {
	file, err := os.Open(`./sample.html`)
	gotorial.OSExitIfError(err)
	node, err := html.Parse(file)
	actual := ElementByID(node, "charset")
	expected := "meta"
	if actual.Data != expected {
		t.Errorf("Expected value is %s but actual is %s", expected, actual.Data)
	}
}
