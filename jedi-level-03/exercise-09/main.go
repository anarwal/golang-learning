package main

import "fmt"

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {
	favSport := "soccer"
	switch favSport {
	case "soccer":
		fmt.Println("Favorite sport is soccer")
	case "baseball":
		fmt.Println("Favorite sport is baseball")
	default:
		fmt.Println("Doesnt like any sports")
	}
}
