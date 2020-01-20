package main

import (
	"fmt"
	"time"
)

// Create a for loop using syntax: `for condition { }` and have it print out the years you have been alive.

func main() {
	i := 1991
	for i <= time.Now().Year() {
		fmt.Println(i)
		i++
	}
}
