package main

import (
	"fmt"
	"sort"
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
}

func main() {
	result := topoSort(prereqs)
	for k, v := range result {
		fmt.Printf("%v: %v\n", v, k)
	}
}

func topoSort(m map[string][]string) map[string]int {
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

//m 受講科目が決定するとValとして、必要科目のスライスが取得される
func topoSortOriginal(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			seen[item] = true
			visitAll((m[item])) //必要科目のさらに必要科目を再帰的に検索
			order = append(order, item)
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
