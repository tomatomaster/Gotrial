package gotorial

import "os"
import "fmt"

//OSExit exit app if err != nil
func OSExit(err error) {
	if err != nil {
		os.Exit(1)
	}
}

//OSExitWithCause print cause and exit app if err != nil
func OSExitWithCause(cause string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v %v", cause, err)
		os.Exit(1)
	}
}
