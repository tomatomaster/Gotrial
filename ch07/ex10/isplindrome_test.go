// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	sample := []struct {
		input palind
		want  bool
	}{
		{"test", false},
		{"aaaa", true},
		{"a", true},
		{" ", true},
		{" yama       amay ", true},
	}

	for _, s := range sample {
		if s.want != IsPalindrome(s.input) {
			t.Fatalf("IsPalindrome(%s) = %v but want %v\n", s.input, IsPalindrome(s.input), s.want)
		}
	}
}
