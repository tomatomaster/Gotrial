// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"./equal"
)

func main() {
	x := 10.0
	y := 9.9999
	r := equal.Equal(x, y)
	fmt.Printf("%v\n", r)

	x = 10.0
	y = 9.999999999999999999999999
	r = equal.Equal(x, y)
	fmt.Printf("%v\n", r)
}
