package main

import "fmt"

// Using a COMPOSITE LITERAL: create an ARRAY which holds 5 VALUES of TYPE int; assign VALUES to each index position; Range over the array and print the values out; Using format printing; print out the TYPE of the array

func main() {
	a := []int{3, 4, 5, 6, 7}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)

}
