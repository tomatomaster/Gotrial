// Copyright © 2017 Ryutarou Ono.

package main

import (
	"testing"

	"net/http"

	"./params"
)

func TestUnpackOK(t *testing.T) {
	urlOk := "http://host?l=test&m=1013214&c=5101234567890123"
	req, _ := http.NewRequest("GET", urlOk, nil)
	if err := searchCopy(nil, req); err != nil {
		t.Errorf("Error %v", err)
	}
}

func TestUnpackMailNG(t *testing.T) {
	urlOk := "http://host?l=test&m=10132145&c=5101234567890123"
	req, _ := http.NewRequest("GET", urlOk, nil)
	if err := searchCopy(nil, req); err == nil {
		t.Errorf("Test NG %v Expected err != nil", err)
	}
}

func TestUnpackCardNG(t *testing.T) {
	urlOk := "http://host?l=test&m=10132145&c=234234"
	req, _ := http.NewRequest("GET", urlOk, nil)
	if err := searchCopy(nil, req); err == nil {
		t.Errorf("Test NG %v Expected err != nil", err)
	}
}

//Copy search method and modify partially to test.
func searchCopy(resp http.ResponseWriter, req *http.Request) error {
	var data struct {
		Labels []string `http:"l"`
		Mail   int      `http:"m mail"`
		Card   int      `http:"c card"`
	}
	return params.Unpack(req, &data)
	// ...rest of handler...
	//fmt.Fprintf(resp, "Search: %+v\n", data)
}
