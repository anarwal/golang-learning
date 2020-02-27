package main

import "fmt"

// Method Sets
// Create a person struct
// Attach a method to speak to type pointer to a person
//   - *person
// Create a type human interface
//   - to implicitly implement the interface, a human must have the speak method
// Create func "saySomething"
//   - have it take in a humain as a parameter
//   - have it call the speak method
// Show the following in your code
//   - you can pass type *person into saySomething
//   - you cannot pass type person into saySomething

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I can talk")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{name: "John"}
	saySomething(&p)
}
