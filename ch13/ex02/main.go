// Copyright © 2017 Ryutarou Ono.

package main

import "fmt"

type link struct {
	value string
	tail  *link
}

func main() {
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c // a->b->a ... c->c->c ...
	result := IsLoop(a)
	fmt.Printf("Result:%v\n", result)
	result = IsLoop(b)
	fmt.Printf("Result:%v\n", result)
	result = IsLoop(c)
	fmt.Printf("Result:%v\n", result)
	d, e, f := &link{value: "d"}, &link{value: "e"}, &link{value: "f"}
	d.tail, e.tail, f.tail = f, d, nil //d->f->nil e->d->f->nil f->nil
	result = IsLoop(d)
	fmt.Printf("Result:%v\n", result)
}
