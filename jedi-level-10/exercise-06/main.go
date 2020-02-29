package main

import "fmt"

// write a program that
//   - puts 100 numbers to a channel
//   - pull the numbers off the channel and print them

func main() {
	c := 100
	chanresponse := sender(c)
	receive(chanresponse)
	fmt.Println("program done")
}

func sender(num int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < num; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for e := range c {
		fmt.Println(e)
	}
}
