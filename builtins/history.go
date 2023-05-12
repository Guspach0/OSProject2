package builtins

import (
	"fmt"
	"io"
)

var HISTORY [1024]string
var SIZE = 0
var err error

func StoreHistory(input string) {
	// reset size and over write if at max size
	if SIZE == 1023 {
		SIZE = 0
	}
	HISTORY[SIZE] = input
	SIZE++
}
func History(w io.Writer, args ...string) error {
	// show history
	if len(args) == 0 {
		for i, v := range HISTORY {
			if i < SIZE {
				fmt.Fprintf(w, "%d\t%s", i+1, v)
			}
		}
	} else if len(args) == 1 && args[0] == "-c" {
		// clear history
		SIZE = 0
		HISTORY = [1024]string{}
		var msg = "cleared history"
		fmt.Fprintln(w, msg)
	} else {
		// display usage
		var error = "usage:\nhistory - list history\nhistory -c - clear history"
		_, err := fmt.Fprintln(w, error)
		return err
	}
	return err
}
