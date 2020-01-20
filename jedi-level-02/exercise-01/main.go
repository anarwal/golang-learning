package main

import "fmt"

// program that prints a number in decimal, binary, and hex

func main() {
	x := 42
	fmt.Printf("%v\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)
}
