package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) != 2 {
		os.Exit(1)
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
		defer fp.Close()
	}
	result := wordfreq(fp)
	for k, v := range result {
		fmt.Printf("%v\t%v\n", k, v)
	}
}

func wordfreq(file *os.File) (m map[string]int) {
	m = make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m[scanner.Text()]++
	}
	return m
}
