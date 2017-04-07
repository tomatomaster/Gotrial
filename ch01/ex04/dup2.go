package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		}
		log.Print("Success to open file")
		countLines(os.Stderr, counts)
		f.Close()
	}
	log.Print("Finish to read all files")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	if input.Scan() {
		log.Print(input.Text())
	}

	for input.Scan() {
		log.Print("test")
		counts[input.Text()]++
	}
	log.Print("Finish read file")
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input:", err)
	}
}
