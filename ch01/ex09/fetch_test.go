package main

import "testing"

import "strings"
import "bytes"
import "net/http"

var f = printStatusCode

func TestStatuCode(t *testing.T) {
	buf := &bytes.Buffer{}
	resp, _ := http.Get("http://google.com")
	f(buf, resp)
	if !strings.Contains(buf.String(), "StatusCode") {
		t.Error("Failed")
	}
}
