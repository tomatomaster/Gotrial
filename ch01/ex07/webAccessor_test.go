package main

import "testing"

func TestGetRespBody(t *testing.T) {
	if !PrintRespBody("http://google.com") {
		t.Error("Fail")
	}
}
