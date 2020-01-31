package main

import (
	"fmt"
	"math"
)

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// 	- circle area= π r 2
//	- square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type rectangle struct {
	length float32
	width  float32
}

type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println("The area -->", s.area())
}

func main() {
	r1 := rectangle{3, 5}
	c1 := circle{5}

	// Printing area directly
	fmt.Println("Area of rectangle is ", r1.area())
	fmt.Println("Area of circle is ", c1.area())

	// Using the function to print area
	info(r1)
	info(c1)
}
