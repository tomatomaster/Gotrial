package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {

	fmt.Printf("Sunday is %d\n", Sunday)
	fmt.Printf("FlagUp is %b\n", FlagUp)
	fmt.Printf("FlagBroadcast is %b\n", FlagBroadcast)
	fmt.Printf("FlagLoopback is %b\n", FlagLoopback)
}
