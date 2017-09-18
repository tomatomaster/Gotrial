// Copyright © 2017 Ryutarou Ono.

package main

import (
	"net/url"
	"testing"
)

func TestPack(t *testing.T) {
	type book struct {
		title string
		num   int
	}
	sample := book{title: "dream",
		num: 10}
	base := "http://bookstore"
	u, _ := url.Parse(base)
	u = Pack(u, &sample)
	if u.String() == base+"?num=10&title=dream" || u.String() == base+"?title=dream&num=10" {
		return
	}
	t.Fatalf("Error %v", u.String())
}
