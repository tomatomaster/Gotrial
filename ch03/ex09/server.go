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
	x, _ := strconv.Atoi(query.Get("x"))
	y, _ := strconv.Atoi(query.Get("y"))
	draw(w, x, y)
}
