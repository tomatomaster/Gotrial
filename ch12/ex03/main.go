// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"log"

	"./sexpr"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	TestComplex     complex64
	TestInterface   interface{}
	TestFloat       float64
}

func main() {
	strangelove := Movie{
		Title:    "Strangelove",
		Subtitle: "How I Learned to Stop Worring and Lova the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":      "Peter Sellers",
			"Gen. Buck Turgidson":  "George C. Scott",
			`maj.T.J."King" Kong"`: "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Director(Nomin.)",
			"Best Director(Nomin.)",
		},
		TestComplex:   1 + 2i,
		TestInterface: []int{1, 2, 3},
		TestFloat:     1.234,
	}
	out, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	for _, outb := range out {
		fmt.Printf("%s", string(outb))
	}
}
