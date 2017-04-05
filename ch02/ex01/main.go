package main

import "./tempconv"
import "fmt"

func main() {
	fmt.Printf("tempconv.CToK converts BoilingC(%v) to %v\n", tempconv.BoilingC, tempconv.CToK(tempconv.BoilingC))
}
