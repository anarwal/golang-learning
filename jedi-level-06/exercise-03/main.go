package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func main() {
	foo()
	defer bar()
	mat()
}

func foo() {
	fmt.Println("1. I am foo")
}

func bar() {
	fmt.Println("2. I am bar")
}

func mat() {
	fmt.Println("3. I am mat")
}
