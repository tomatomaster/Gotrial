// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"log"

	"encoding/json"

	"./sexpr"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
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
	}
	fmt.Println("Original Json Marshal")
	jsonout, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(jsonout))
	fmt.Println("Json Unmarshal")
	var movie Movie
	err = json.Unmarshal(jsonout, &movie)
	if err != nil {
		log.Printf("error %+v", err)
	}
	fmt.Printf("%v\n", movie)
}
