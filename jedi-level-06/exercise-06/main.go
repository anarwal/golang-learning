package main

import "fmt"

// Build and use an anonymous func

func main() {
	func() {
		fmt.Println("This function doesn't have a name...")
	}()

	fmt.Println("Function block ended!")
}
