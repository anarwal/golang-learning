package main

import "fmt"

// Building on the jedi-level-03, exercise-06, create a program that uses “else if” and “else”.

func main() {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("This number: %v is divisible by 2\n", i)
		} else if i%3 == 0 {
			fmt.Printf("This number: %v is divisible by 3\n", i)
		} else {
			fmt.Printf("This number: %v is not divisible by 2 or 3\n", i)
		}
	}
}
