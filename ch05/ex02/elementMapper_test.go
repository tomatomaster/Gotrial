package main

import (
	"os"
	"testing"

	"github.com/tomatomaster/gotorial/gotorial"
	"golang.org/x/net/html"
)

var f = mapper

func TestMapper(t *testing.T) {
	file, err := os.Open(`./sample.html`)
	gotorial.OSExitIfError(err)
	node, err := html.Parse(file)
	gotorial.OSExitIfError(err)
	actual := make(map[string]int)
	actual = f(actual, node)
	expected := map[string]int{"html": 1, "head": 1, "body": 1, "a": 3}
	if len(actual) != len(expected) {
		t.Errorf("actual len is %v expected len is %v\n", actual, len(expected))
	}

	for k, v := range actual {
		if expected[k] != v {
			t.Errorf("key %v actual value is %v expected value is %v\n", k, v, expected[k])
		}
	}

}
