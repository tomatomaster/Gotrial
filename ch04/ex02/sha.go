package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Can not find any input")
		fmt.Println("Usage: sha [any input]")
		os.Exit(0)
	}
	argVal := os.Args[1]
	sha(argVal)
}

func sha(argVal string) {
	shaType := flag.Int("t", 256, "Select 256 384 512")
	flag.Parse()
	argByte := []byte(argVal)
	if *shaType == 256 {
		fmt.Printf("sha256: %x\n", sha256.Sum256(argByte))
	} else if *shaType == 384 {
		fmt.Printf("sha384: %x\n", sha512.Sum384(argByte))
	} else if *shaType == 512 {
		fmt.Printf("sha512: %x\n", sha512.Sum512(argByte))
	} else {
		fmt.Println("Select 256, 384 or 512")
	}
}
