// Copyright © 2017 Ryutarou Ono.

package main

import "fmt"

func main() {
	type Cycle struct {
		Value int
		tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	fmt.Println("First Call Start")
	Display("cycle", c)
	fmt.Println("First Call End")
	fmt.Println("Second Call Start")
	Display("cycle", c)
	fmt.Println("Second Call End")
}
