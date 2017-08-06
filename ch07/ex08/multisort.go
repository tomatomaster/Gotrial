// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"sort"
)

type Human struct {
	Name   string
	Height int32
	Weight int32
}

var ClassMates = []*Human{
	{"Tarou", 160, 48},
	{"Kyouko", 158, 50},
	{"Katou", 172, 58},
}

type byHeight []*Human

func (h byHeight) Len() int {
	return len(h)
}

func (h byHeight) Less(i, j int) bool {
	return h[i].Height < h[j].Height
}

func (h byHeight) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	sort.Sort(sort.Reverse(byHeight(ClassMates)))
	for _, h := range ClassMates {
		fmt.Printf("%v\n", h)
	}
}
