package main

import "fmt"

// get following code working: https://play.golang.org/p/j-EA6003P0

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

// User buffered channel
//
// func main() {
// 	c := make(chan int,1)
// 	c <-42
// 	fmt.Println(<-c)
// }
