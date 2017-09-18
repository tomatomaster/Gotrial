// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"

	"strings"

	"./sexpr"
)

type Test struct {
	Test string `sexpr:"hogehoge"`
}

func TestUnmarshal(t *testing.T) {
	sample := Test{
		Test: "test",
	}
	out, err := sexpr.Marshal(sample)
	if err != nil {
		t.Fatalf("Unexpected error %v\n", err)
	}
	if !strings.Contains(string(out), "hogehoge") {
		t.Errorf("%s", string(out))
	}
	var unmarshal Test
	sexpr.Unmarshal(out, &unmarshal)
	if unmarshal.Test != "test" {
		t.Errorf("%s", unmarshal.Test)
	}
}
