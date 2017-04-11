package main

import (
	"log"
	"net/http"
	"sync"

	"net/url"
	"strconv"

	"./anime"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	i := getCyclesFrom(r)
	anime.Lissajous(w, i)
}

func getCyclesFrom(r *http.Request) int {
	query := r.URL.RawQuery
	v, _ := url.ParseQuery(query)
	s := v.Get("cycles")
	i, _ := strconv.Atoi(s)
	return i
}
