package main

import "fmt"

// Using the code from exercise jedi-level09/exercise-08, add a record to your map. Now print the map out using the “range” loop

func main() {
	m := make(map[string][]string)

	m["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	m["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for i, v := range m {
		fmt.Println("Key: ", i)
		for j, val := range v {
			fmt.Printf("%v : %v\n", j, val)
		}
	}
	// adding new value
	m["new_mr_bond"] = []string{`not smart`, `too lazy`}

	for i, v := range m {
		fmt.Println("Key: ", i)
		for j, val := range v {
			fmt.Printf("%v : %v\n", j, val)
		}
	}
}
