package main

import "fmt"

// create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

func main() {
	fmt.Printf("India scored %v\n", foo())
	x, y := bar()
	fmt.Printf("England scored %v\n", x)
	fmt.Printf("And the winner was %v\n", y)
}

func foo() int {
	return 300
}

func bar() (x int, y string) {
	x = 200
	y = "India"
	return
}
