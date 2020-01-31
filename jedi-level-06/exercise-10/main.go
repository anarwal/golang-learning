package main

import "fmt"

// Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:

func main() {
	multiple := foo()
	fmt.Println(multiple())
	fmt.Println(multiple())
	fmt.Println(multiple())
	fmt.Println(multiple())
}

func foo() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}
