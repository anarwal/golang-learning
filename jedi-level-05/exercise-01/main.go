package main

import "fmt"

// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

type person struct {
	firstName         string
	lastName          string
	favIceCreamFlavor []string
}

func main() {
	p1 := person{}
	p1.firstName = "john"
	p1.lastName = "doe"
	p1.favIceCreamFlavor = []string{"vanilla", "chocolate"}

	p2 := person{"sam", "smarty", []string{"chocolate", "black-current"}}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favIceCreamFlavor {
		fmt.Println(i, v)
	}

	for i, v := range p2.favIceCreamFlavor {
		fmt.Println(i, v)
	}

}
