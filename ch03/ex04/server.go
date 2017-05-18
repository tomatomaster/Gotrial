package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	width, _ := strconv.Atoi(query.Get("with"))
	height, _ := strconv.Atoi(query.Get("height"))
	w.Header().Set("Content-Type", "image/svg+xml")
	Write(w, width, height, "#ff0000")
}
