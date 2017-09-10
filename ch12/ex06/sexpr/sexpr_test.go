// Copyright © 2017 Ryutarou Ono.

package sexpr

import (
	"log"
	"testing"

	"strings"

	"gopl.io/ch12/sexpr"
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

func TestMarshal(t *testing.T) {
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
		if strings.Contains(string(outb), "Oscars ") {
			t.Errorf("Illegal %v", string(outb))
		}
	}

}
