package recursion

// factorial using loops

import (
	"fmt"
)

func main() {
	n := factorial(5)
	fmt.Println(n)
}

func factorial(n int) int {
	y := 1
	for ; n > 0; n-- {
		y *= n
	}
	return y
}
