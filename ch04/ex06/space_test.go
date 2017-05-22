package main

import (
	"testing"
)

var f = compress

func Test1(t *testing.T) {
	sample := "今日は  世界"
	expected := "今日は 世界"
	result := f([]byte(sample))
	actual := string(result)
	if actual != expected {
		t.Errorf("Expected: %s Actual %s", expected, actual)
	}
}

func Test2(t *testing.T) {
	sample := "　今日は        世 　　 あ     界"
	expected := "今日は 世 界"
	result := f([]byte(sample))
	actual := string(result)
	if actual != expected {
		t.Errorf("Expected: %s Actual %s", expected, actual)
	}
}
