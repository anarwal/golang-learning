package main

import (
	"fmt"
	"time"
)

// Create a for loop using syntax: `for { }` and have it print out the years you have been alive.

func main() {
	i := 1991
	for {
		if i > time.Now().Year() {
			break
		}
		fmt.Println(i)
		i++
	}
}
