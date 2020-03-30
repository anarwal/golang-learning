package main

import (
	"fmt"
)

// Create a struct “customErr” which implements the builtin.error interface.
// Create a func “foo” that has a value of type error as a parameter.
// Create a value of type “customErr” and pass it into “foo”

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("An error as occuredm which gave value: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "You blew it away",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(" foo gave -", e)
}
