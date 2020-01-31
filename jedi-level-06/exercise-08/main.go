package main

import "fmt"

// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

func main() {
	x := foo()
	fmt.Println(x())
}

func foo() func() int {
	return func() int {
		sum := 0
		for x := 0; x <= 10; x++ {
			sum += x
		}
		return sum
	}
}
