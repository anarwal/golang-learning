package main

// Unmarshal given json to array

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	jsonData := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27}]`

	var u []user

	err := json.Unmarshal([]byte(jsonData), &u)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
