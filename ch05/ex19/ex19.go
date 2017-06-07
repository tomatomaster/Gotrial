package main

import "fmt"

func main() {
	fmt.Println(*magic())
}

func magic() (p *interface{}) {
	p = new(interface{})
	defer func() {
		r := recover()
		if r != nil {
			*p = r
		}
	}()
	panic("panic panic!")
}
