// Copyright © 2017 Ryutarou Ono.

package main

import (
	"os"
	"testing"
)

func TestCheckDependency(t *testing.T) {
	os.Args = []string{"", "unsafe"}
	//正解が分からない...
	CheckDependency()
}
