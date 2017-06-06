package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input string
		expec string
		repla string
	}{
		{input: "abcdefooghijk", expec: "abcdefghijk", repla: "f"},
		{input: "", expec: "", repla: "f"},
		{input: "a", expec: "a", repla: "f"},
		{input: "Whatfoofo", expec: "Whathellofo", repla: "hello"},
	}
	for _, test := range tests {
		actual := expand(test.input, func(s string) string {
			return test.repla
		})
		if actual != test.expec {
			t.Fatalf("Expected: %v But %v", test.expec, actual)
		}
	}

}
