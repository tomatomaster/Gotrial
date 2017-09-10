package main

import (
	"testing"
)

func TestBit(t *testing.T) {
	tests := []struct {
		input  int
		want   int
		result bool
	}{
		{input: 1, want: 2, result: false},
		{input: 1, want: 1, result: true},
		{input: 4, want: 4, result: true},
		{input: 10000, want: 200000, result: false},
	}

	var set IntSet
	for _, test := range tests {
		set.Add(test.input)
		actual := set.Has(test.want)
		if actual != test.result {
			t.Fatalf("Expected %v but actual %v\n", test.result, actual)
		}
	}
}

func Test(t *testing.T) {
	tests := []struct {
		input   []int
		del     []int
		has     int
		wantLen int
		result  bool
	}{
		{input: []int{1, 2, 3, 4}, del: []int{2, 3}, has: 2, wantLen: 4, result: false},
		{input: []int{100}, del: []int{100, 3}, has: 2, wantLen: 1, result: false},
		{input: []int{100, 111, 12314, 1231}, del: []int{111, 3}, has: 100, wantLen: 4, result: true},
	}

	for _, test := range tests {
		var set IntSet
		for i := range test.input {
			set.Add(i)
		}
		len := set.Len()
		if len != test.wantLen {
			t.Fatalf("Len Error len %v want %v", len, test.wantLen)
		}
		for i := range test.del {
			set.Remove(i)
		}
		if set.Has(test.has) != test.result {
		}
	}
}

func TestAdd(t *testing.T) {
	var set IntSet
	var m map[int]bool
	set.Add(1)
	m[1] = true
	set.Has(1)

}
