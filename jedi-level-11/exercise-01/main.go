package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// starting with https://play.golang.org/p/3W69TH4nON, instead of using the blank identifier, make sure the code is checking and handling the error.

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		// using simple error handling
		// fmt.Println(err)
		// we will use "Fatal" as if the function can not convert to json it is of no use
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
