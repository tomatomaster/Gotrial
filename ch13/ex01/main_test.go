// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"

	"./equal"
)

func TestEqual(t *testing.T) {
	samples := []struct {
		x, y float64
		want bool
	}{
		{x: 10.0, y: 9.0, want: false},
		{x: 9.0, y: 9.00000000000000001, want: true},
		{x: 9.0, y: 8.99999999999999999, want: true},
	}

	for _, s := range samples {
		r := equal.Equal(s.x, s.y)
		if r != s.want {
			t.Fatalf("Input x:%f y:%f actual:%v want:%v", s.x, s.y, r, s.want)
		}
	}

}
