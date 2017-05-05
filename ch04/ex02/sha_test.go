package main

import (
	"flag"
	"testing"
)

var f = sha

func TestSha(t *testing.T) {
	flag.Set("t", "512")
	f("t")
}
