package main

// Starting with: https://play.golang.org/p/MvL6uamrJP, pull the values off the channel using a select statement
import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case m1 := <-c:
			fmt.Println("received", m1)
		case <-q:
			return
		}
	}

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
