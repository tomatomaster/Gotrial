// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"bytes"

	"./sexpr"
)

type Data struct {
	Title, SubTitle string
	Year            int
}

//Give up...
func main() {

	var data Data
	input := `((Title "test") (SubTitle "Fxxk you") (Year 123))`
	reader := bytes.NewReader([]byte(input))
	decoder := sexpr.NewReader(reader)
	decoder.UnmarshalReader(&data)

	fmt.Printf("%v", data)
}
