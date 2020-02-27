package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Using goroutines, create an incrementer program
//   - have a variable to hold the incrementer value
//   - launch a bunch of goroutines
//      * each goroutine should
//         # read the incrementer value
//             -- store it in a new variable
//         # yield the processor with runtime.Gosched()
//         # increment the new variable
//         # write the value in the new variable back to the incrementer variable
//   - use waitgroups to wait for all of your goroutines to finish
//   - the above will create a race condition.
//   - Prove that it is a race condition by using the -race flag

func main() {
	incrementer := 0
	const c = 50
	var wg sync.WaitGroup
	wg.Add(c)

	for i := 0; i < c; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Incrementer: ", incrementer)
}
