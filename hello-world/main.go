package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is my first golang program!")
	foo()
	fmt.Println("here is some more intro")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am an extra function")
}
