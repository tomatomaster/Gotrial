// Copyright © 2017 Ryutarou Ono.

package main

import (
	"flag"
	"testing"

	"./tempconv"
)

func TestFlag(t *testing.T) {
	var temp = tempconv.CelsiusFlag("temp", 0, "the temperature")

	cases := []struct {
		in       string
		expected tempconv.Celsius
	}{
		{"0K", tempconv.Celsius(-273.5)},
		{"-1K", tempconv.Celsius(-273.5)},
		{"100K", tempconv.Celsius(-173.5)},
		{"273.5K", tempconv.Celsius(0)},
	}

	for _, c := range cases {
		flag.Set("temp", c.in)
		flag.Parse()
		if c.expected != *temp {
			t.Fatalf("Expected %v Actual %v", c.expected, *temp)
		}
	}

}
