package main

import "fmt"

// Building on the code from the exercise-04, at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// In func main, now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE, then use the “=” operator to ASSIGN that value to “y”, print out the value stored in “y”, print out the type of “y”

type hakunamata int

var x hakunamata
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
