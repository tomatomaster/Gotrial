// Copyright © 2017 Ryutarou Ono.

package main

import (
	"bufio"
	"os"
	"testing"
)

func TestLimit(t *testing.T) {
	file, _ := os.Open("./test")
	lr := LimitReader(file, 8)
	scanner := bufio.NewScanner(lr)
	for scanner.Scan() {
		actual := scanner.Text()
		if actual != "12345678" {
			t.Errorf("Error actual %s expect %s", actual, "12345678")
		}
	}
}
