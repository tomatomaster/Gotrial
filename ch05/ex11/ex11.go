package main

import (
	"fmt"
	"log"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func main() {
	result := topoSort(prereqs)
	for k, v := range result {
		fmt.Printf("%v: %v\n", v, k)
	}
}

func topoSort(m map[string][]string) map[string]int {
	checkcirculation(m)
	order := make(map[string]int)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	count := 0
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				count++
				order[item] = count
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

func checkcirculation(m map[string][]string) {
	for fk, fv := range m { //全ての要素に対して、
		for _, values := range fv { //前提科目が受講科目を前提科目としていないか確認する
			for _, sv := range m[values] {
				if fk == sv {
					log.Fatalf("Find Circulatio Dependency!! \n \"%s\"  \"%s\" ", fk, values)
				}
			}
		}
	}
}
