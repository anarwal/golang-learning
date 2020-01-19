package main

import "fmt"

const y = 43

//DECLARE there is a var with identifier z
//variable z is of type "int"
// Variables declared without an explicit initial value are given their zero value.
// The zero value is: 0 for numeric types, false for the boolean type, and "" (the empty string) for strings.
var z int

func main() {
	//short declaration
	//DECLARE a variable and ASSIGN A VALUE (of certain type)

	x := 42
	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("again: ", y)
}
