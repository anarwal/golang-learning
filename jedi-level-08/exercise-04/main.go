package main

// Sort the []int and []string for each person

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 43, 2, 16, 1, 23, 12, 1, 5}
	xs := []string{"random", "rainbow", "randy", "snow", "ball"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
