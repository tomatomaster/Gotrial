// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gopl.io/ch7/eval"
)

func main() {
	http.HandleFunc("/", input)
	http.HandleFunc("/input", cal)
	http.ListenAndServe(":8000", nil)
}

func cal(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("input")
	if val == "" {
		return
	}
	expr, err := eval.Parse(val)
	if err != nil {
		log.Fatal(err)
	}
	result := expr.Eval(eval.Env{})
	fmt.Fprintf(w, "%f", result)
}

func input(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("input.gtpl")
	t.Execute(w, nil)
}
