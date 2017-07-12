// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"math"
	"time"
)

const GOROUTINE = 100

var result = make(chan float64)

func main() {
	tS := time.Now()
	for i := 0; i < GOROUTINE; i++ {
		go splitFunction(i*1000000, (i+1)*1000000-1)
	}
	var sum float64

	for i := 0; i < GOROUTINE; i++ {
		sum += <-result
	}
	fmt.Printf("SUM: %f\n", sum)
	fmt.Printf("Range: %v\n", time.Since(tS))
	tS = time.Now()
	highcostFunction()
	fmt.Printf("Range: %v\n", time.Since(tS))
}

func splitFunction(start, end int) {
	var sum float64
	for i := start; i <= end; i++ {
		sum += math.Gamma(10) + math.Atan(10) + math.Sin(float64(10))
	}
	result <- sum
}

func highcostFunction() {
	var sum float64
	for i := 0; i < 100*1000000; i++ {
		sum += math.Gamma(10) + math.Atan(10) + math.Sin(float64(10))
	}
	fmt.Printf("SUM: %f \n", sum)
}
