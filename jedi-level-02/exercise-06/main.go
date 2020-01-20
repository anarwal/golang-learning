package main

import (
	"fmt"
	"time"
)

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

var currentYear = time.Now().Year()

const (
	a = 2020
	b = a + iota
	c = a + iota
	d = a + iota
	e = a + iota
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
