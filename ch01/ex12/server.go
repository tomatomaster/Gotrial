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
	query := r.URL.RawQuery
	v, _ := url.ParseQuery(query)
	s := v.Get("cycles")
	i, _ := strconv.Atoi(s)
	anime.Lissajous(w, i)
}
