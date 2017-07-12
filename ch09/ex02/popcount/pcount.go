package popcount

import (
	"fmt"
	"sync"
)

var pc [256]byte
var initPopcountOnce sync.Once

func initTable() {
	fmt.Println("初期化するよ！")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount return the number of set bits of x
func PopCount(x uint64) int {
	initPopcountOnce.Do(initTable)
	var count int
	var i uint
	for i = 0; i <= 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//PopCountOriginal return the number of set bits of x
func PopCountOriginal(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
