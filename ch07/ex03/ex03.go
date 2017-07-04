package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
	string      fmt.Stringer
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func (m map[string]int) modify() {

}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	var t *tree
	t = add(t, 100)
	t = add(t, 120)
	t = add(t, 1)
	t = add(t, 13)
	t = add(t, 11231)
	t = add(t, 131)
	fmt.Println(t)
}

func (t *tree) String() string {
	return childValue(t)
}

func childValue(t *tree) string {
	if t == nil {
		return ""
	}
	top := t.value
	var l int
	var r int
	var result string
	if t.left != nil {
		l = t.left.value
		result += childValue(t.left)
	}
	if t.right != nil {
		r = t.right.value
		result += childValue(t.right)
	}
	result += fmt.Sprintf("{%d %d %d}\n", l, top, r)
	return result
}
