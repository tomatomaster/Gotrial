package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fPath := os.Args[1]
	file, err := os.Open(fPath)
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	var command string
	for sc.Scan() {
		line := sc.Text()
		url := getURL(line)
		command += "https://" + url + " "
	}
	fmt.Print(command)
}

func getURL(s string) string {
	s = strings.Split(s, ",")[1]
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "/\"")
	return s
}
