// Copyright © 2017 Ryutarou Ono.
package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestParse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"<html><title>test</title></html>", "html"},
		{"<><title>test</title></>", ""},
		{"<KOME><title>test</title></KOME>", "KOME"},
	}

	for _, c := range cases {
		node, _ := Parse(c.in)
		if node.Type != html.ElementNode {
			continue
		}
		actual := node.Data
		if actual != c.want {
			t.Errorf("Actual is %s Want is %s \n", actual, c.want)
		}
	}

}
