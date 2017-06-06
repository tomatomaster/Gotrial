package main

import (
	"log"
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		inputs []int
	}{
		{inputs: []int{1, 2, 3, 4, 5, 6, 100}},
		{inputs: []int{100, 131314124, 31231, 1313213, 213231231, 123123131, 123131231}},
	}
	for _, test := range tests {
		var set IntSet
		set.AddAll(test.inputs...)
		for _, has := range test.inputs {
			if !set.Has(has) {
				log.Fatalf("Set does't have %v \n", has)
			}
		}
	}
}
