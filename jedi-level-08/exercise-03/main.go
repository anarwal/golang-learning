package main

// encode to JSON sending results to stdout

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is not guarantee of innovation",
		},
	}

	err := json.NewEncoder(os.Stdout).Encode(u1)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
