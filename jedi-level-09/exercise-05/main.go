package main

// Fix the race condition you created in exercise #4 by using package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const c = 50
	var wg sync.WaitGroup
	var incrementer int64

	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Incrementer: ", incrementer)
}
