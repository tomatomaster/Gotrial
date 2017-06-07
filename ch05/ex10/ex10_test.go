package main

import (
	"testing"
)

func TestEx10(t *testing.T) {
	prereqs := map[string][]string{
		"algorithms": {"data structures"},

		"data structures": {"discrete math"},
		"discrete math":   {"intro to programming"},
	}

	correct := map[string]int{
		"intro to programming": 1,
		"discrete math":        2,
		"data structures":      3,
		"algorithms":           4,
	}

	result := topoSort(prereqs)
	for key, index := range result {
		if correct[key] != index {
			t.Fatalf("Expected: %v, But: %v", correct[key], index)
		}
	}
}
