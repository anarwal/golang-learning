package main

import "fmt"

// Create a variable of type string using a raw string literal. Print it.

func main() {
	var x string = `My name is Bond, 
		"James Bond"`
	fmt.Printf("%v\n", x)
}
