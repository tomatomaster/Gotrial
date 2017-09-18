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
	IYear           complex128
	Bool            bool
	Float           float64
	Inter           interface{}
}

func main() {

	var data Data
	input := `((Title "test") (SubTitle "Fxxk you") (Year 123) (IYear #C(1.0 2.0)) (Bool nil) (Float 1.23) (Inter ([2]int (0 2))))`
	reader := bytes.NewReader([]byte(input))
	//fmt.Println("Data type Title string, Subtitle string, Year int")
	//fmt.Println("Ctrl + G")
	//reader := bufio.NewReader(os.Stdin)
	decoder := sexpr.NewReader(reader)
	decoder.UnmarshalReader(&data)

	fmt.Printf("%v\n", data)
}
