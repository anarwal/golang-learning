package main

import "fmt"

// Assign a func to a variable, then call that func

func main() {
	f := func() {
		fmt.Println("This function doesn't have a name and is assigned to a var...")
	}

	f()
	fmt.Println("Function block ended!")
}
