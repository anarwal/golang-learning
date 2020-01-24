package main

import "fmt"

// Using the code from exercise jedi-level09/exercise-09, delete a record from your map. Now print the map out using the “range” loop

func main() {
	m := make(map[string][]string)

	m["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	fmt.Println("Printing for the first time")
	for i, v := range m {
		fmt.Println("Key: ", i)
		for j, val := range v {
			fmt.Printf("%v : %v\n", j, val)
		}
	}
	// adding new value
	m["new_mr_bond"] = []string{`not smart`, `too lazy`}

	fmt.Println("Printing after addition")
	for i, v := range m {
		fmt.Println("Key: ", i)
		for j, val := range v {
			fmt.Printf("%v : %v\n", j, val)
		}
	}

	//deleting value
	delete(m, "new_mr_bond")

	fmt.Println("Printing after deletion")
	for i, v := range m {
		fmt.Println("Key: ", i)
		for j, val := range v {
			fmt.Printf("%v : %v\n", j, val)
		}
	}
}
