package main

import "os"

func ExampleEcho() {
	os.Args = []string{"test", "test"}
	echo()
	// Output:
	// test test
}

func ExampleFailEcho() {
	os.Args = []string{"something", "toDo"}
	echo()
	// Output:
	// something toDo
}
