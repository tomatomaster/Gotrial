// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"
)

func TestDisplay(t *testing.T) {
	type Cycle struct {
		Value int
		tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("cycle", c)
}
