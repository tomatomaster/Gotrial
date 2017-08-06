// Copyright © 2017 Ryutarou Ono.

package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		sample  string
		sep     string
		wantVal []string
		wantLen int
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}, 3},
		{"a,b,c", ",", []string{"a", "b", "c"}, 3},
		{"a\"b\"c", "\"", []string{"a", "b", "c"}, 3},
	}

	for _, test := range tests {
		words := strings.Split(test.sample, test.sep)
		if got := len(words); got != test.wantLen {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.sample, test.sep, got, test.wantLen)
		}
	}
}
