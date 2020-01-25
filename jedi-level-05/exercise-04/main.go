package main

import "fmt"

// Create and use an anonymous struct.

func main() {
	lang := struct {
		name  string
		isOOP bool
	}{
		"golang",
		true,
	}
	fmt.Println(lang)
}
