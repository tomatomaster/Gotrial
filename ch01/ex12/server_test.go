package main

import (
	"net/http"
	"net/url"
	"testing"
)

var testF = getCyclesFrom

func TestGetCyclesFrom(t *testing.T) {
	r := new(http.Request)
	r.URL = new(url.URL)
	r.URL.RawQuery = "cycles=2"
	v := testF(r)
	if v != 2 {
		t.Error("Failed")
	}
}
