package main

import "fmt"

// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice
// "James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//Range over the records, then range over the data in each record.

func main() {
	records := [][]string{}

	row1 := []string{"James", "Bond", "Shaken, not stirred"}
	row2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	records = append(records, row1)
	records = append(records, row2)

	for i, v := range records {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("At position %v value %v \n", j, val)
		}
	}
}
