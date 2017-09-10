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
		Subtitle: "How I Learned to Stop Worring and Lova the Bomb",
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":      "Peter Sellers",
			"Gen. Buck Turgidson":  "George C. Scott",
			`maj.T.J."King" Kong"`: "Slim Pickens",
		},
	}
	out, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	for _, outb := range out {
		fmt.Printf("%s", string(outb))
	}
}
