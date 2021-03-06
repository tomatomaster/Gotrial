package main

import "fmt"

const (
	//KB value is 1000
	KB = 1000
	//MB value is 1000000
	MB = 1000000 //KB * KB
	//GB value is 1000000000
	GB = 1000000000 //MB * KB
	//TB value is 1000000000000
	TB = 1000000000000
	//PB value is 1000000000000000
	PB = 1000000000000000
	//EB value is 1000000000000000000
	EB = 1000000000000000000
	//ZB value is 1000000000000000000000
	ZB = 1000000000000000000000
	//YB value is 1000000000000000000000000
	YB = 1000000000000000000000000
)

func main() {
	fmt.Printf("TB value is %d\n", TB)
}
