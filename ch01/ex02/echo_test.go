package main

import "os"

func ExampleEcho() {
	os.Args = []string{"none", "test", "test"}
	main()
	// Output:
	// 1:test
	// 2:test
}

func ExampleFailEcho() {
	os.Args = []string{"none", "something", "toDo"}
	main()
	// Output:
	// 1:something
	// 2:toDo
}
