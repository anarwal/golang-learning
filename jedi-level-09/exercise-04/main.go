package main

import (
	"fmt"
	"sync"
)

// Fix the race condition you created in the previous exercise by using a mutex
//   - it makes sense to remove runtime.Gosched()

func main() {
	incrementer := 0
	const c = 50
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(c)

	for i := 0; i < c; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Incrementer: ", incrementer)
}
