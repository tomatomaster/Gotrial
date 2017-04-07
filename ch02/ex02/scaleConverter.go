package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var arg string

func main() {
	if len(os.Args) == 1 {
		if sc.Scan() {
			arg = sc.Text()
			floatArg, _ := strconv.ParseFloat(arg, 64)
			fmt.Println(p2K(floatArg))
		}
	}
	for _, arg = range os.Args[1:] {
		floatArg, _ := strconv.ParseFloat(arg, 64)
		fmt.Println(p2K(floatArg))
	}
}

func p2K(arg float64) float64 {
	return arg * 0.453592
}
