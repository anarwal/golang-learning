package main

import "fmt"

// Take the code from exercise jedi-level-04/exercise-01, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

type person struct {
	firstName         string
	lastName          string
	favIceCreamFlavor []string
}

func main() {
	p1 := person{"john", "doe", []string{"vanilla", "chocolate"}}
	p2 := person{"sam", "smarty", []string{"chocolate", "black-current"}}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIceCreamFlavor {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
