package main

import (
	"testing"
)

func TestCorner(t *testing.T) {
	i, j := corner(10, 10)
	if i == 0 || j == 0 {
		t.Fatalf("Can't get i j val %f %f", i, j)
	}
}
