package main

import (
	"testing"
)

type sample struct {
	data     []float64
	expected float64
}

func TestMax(t *testing.T) {
	samples := []sample{
		{data: []float64{1, 2, 3},
			expected: 3},
		{data: []float64{10, -10, 100},
			expected: 100},
	}
	for _, sample := range samples {
		m, _ := max(sample.data...)
		if m != sample.expected {
			t.Fatalf("Expected value is %v Actual is %v", sample.expected, m)
		}
	}
}

func TestMin(t *testing.T) {
	samples := []sample{
		{data: []float64{1, 2, 3},
			expected: 1},
		{data: []float64{10, -10, 100},
			expected: -10},
	}
	for _, sample := range samples {
		m, _ := min(sample.data...)
		if m != sample.expected {
			t.Fatalf("Expected value is %v Actual is %v", sample.expected, m)
		}
	}
}
