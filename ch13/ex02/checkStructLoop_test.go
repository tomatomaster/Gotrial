// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"
)

func TestIsLoop(t *testing.T) {
	pA := &link{value: "a"}
	pB := &link{value: "b"}
	pC := &link{value: "c"}
	pD := &link{value: "d"}
	pE := &link{value: "e"}
	pA.tail, pB.tail, pC.tail, pD.tail = pB, pA, pC, pE

	samples := []struct {
		p    interface{}
		want bool
	}{
		{p: pA, want: true},
		{p: pB, want: true},
		{p: pC, want: true},
		{p: pD, want: false},
	}

	for _, s := range samples {
		if IsLoop(s.p) != s.want {
			t.Logf("%v is illegal state.", s)
		}
	}
}
