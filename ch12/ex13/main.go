// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"log"

	"./sexpr"
)

type Movie struct {
	Title    string `sexpr:"Ex13TagName"`
	Subtitle string
	Year     int
	Color    bool
	Actor    map[string]string
	Oscars   []string
	Sequel   *string
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
	out, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	for _, outb := range out {
		fmt.Printf("%s", string(outb))
	}
	fmt.Println()
	var unmarshal Movie
	err = sexpr.Unmarshal(out, &unmarshal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", unmarshal)
}
