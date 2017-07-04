// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestCrawl(t *testing.T) {
	os.Args[1] = "https://github.com/tomatomaster"
	flag.CommandLine.Set("depth", "2")
	main()
	scanner := bufio.NewScanner(os.Stdout)
	for scanner.Scan() {
		fmt.Println(scanner)
	}
}
