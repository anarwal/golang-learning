package main

import "fmt"

// A “callback” is when we pass a func into a func as an argument. For this exercise,
// 	- pass a func into a func as an argument

func main() {

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo(f func(x []int) int, y []int) int {
	m := f(y)
	m++
	return m
}
