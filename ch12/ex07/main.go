// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"

	"bytes"

	"log"

	"./sexpr"
)

type Data struct {
	Title, SubTitle string
	Year            int
	IYear           complex128
}

func main() {
	var data Data
	input := `((Title "test") (SubTitle "Fk you") (Year 123) (IYear #C(1.0 2.0)))`
	reader := bytes.NewReader([]byte(input))
	//fmt.Println("Data type Title string, Subtitle string, Year int")
	//fmt.Println("Ctrl + G")
	//reader := bufio.NewReader(os.Stdin)
	if err := sexpr.NewDecoder(reader).Decode(&data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)
}
