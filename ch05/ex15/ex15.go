package main

import (
	"fmt"
	"math"
)

func main() {
	r1, _ := max(-1, 10, 100, 10)
	r2 := maxv2(-1)
	r3, _ := min(-10, 100, 1)
	fmt.Printf("%v %v %v", r1, r2, r3)
}

//引数無しを許可する
func max(vals ...float64) (max float64, ok bool) {
	if len(vals) == 0 {
		return
	}
	ok = true
	temp := float64(math.MinInt64)
	for _, val := range vals {
		max = math.Max(temp, val)
		temp = max
	}
	return
}

//引数無しを許可しない
func maxv2(val float64, vals ...float64) (max float64) {
	if len(vals) == 0 {
		return val
	}
	temp := float64(math.MinInt64)
	for _, val := range vals {
		max = math.Max(temp, val)
		temp = max
	}
	return max
}

func min(vals ...float64) (min float64, ok bool) {
	if len(vals) == 0 {
		return
	}
	ok = true
	temp := float64(math.MaxFloat64)
	for _, val := range vals {
		min = math.Min(temp, val)
		temp = min
	}
	return
}
