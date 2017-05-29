package main

import "sort"

func topoSort(m map[string][]string) []string {
	var order map[string]string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				visitAll(m[item])
				order
			}
		}
	}
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
}
