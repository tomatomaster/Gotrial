package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestElems(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 100}, expected: []int{1, 100}},
		{input: []int{}, expected: []int{}},
		{input: []int{1, 100, 12313, 121, 31422141, 31231, 3131231, 1213213, 1231312, 31221, 2312132, 1213231312, 31, 1, 1, 1}, expected: []int{1, 100, 12313, 121, 31422141, 31231, 3131231, 1213213, 1231312, 31221, 2312132, 1213231312, 31}},
	}

	for _, test := range tests {
		var set IntSet
		set.AddAll(test.input...)
		sort.Sort(sort.IntSlice(test.expected))
		if !reflect.DeepEqual(test.expected, set.Elems()) {
			t.Fatalf("Test %v Elems %v\n", test.expected, set.Elems())
		}
	}

}
