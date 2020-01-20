package main

import "fmt"

// Create a program that shows the “if statement” in action.

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
