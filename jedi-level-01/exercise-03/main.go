package main

import "fmt"

// Declare three variables using var and provide them package level scope. Provide following type to each variable: x type int, y type string and z type bool.
// Assign x=42, y="james Bond", z = true
// In main func, use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”.
// Print out the value stored by variable “s”

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v \t %v\t %v \t", x, y, z)
	fmt.Println(s)
}
