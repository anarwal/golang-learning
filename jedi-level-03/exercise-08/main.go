package main

import "fmt"

// Create a program that uses a switch statement with no switch expression specified.

func main() {
	switch {
	case true:
		fmt.Println("This is true")
	case false:
		fmt.Println("This statement is not true")
	}
}
