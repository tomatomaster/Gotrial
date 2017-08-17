// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

type database map[string]dollars

func (db database) update(w http.ResponseWriter, r *http.Request) {
	i := r.URL.Query().Get("item")
	p := r.URL.Query().Get("price")
	floatP, err := strconv.ParseFloat(p, 32)
	if err != nil {
		log.Fatal(err)
	}
	db[i] = dollars(floatP)
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	i := r.URL.Query().Get("item")
	delete(db, i)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
