package main

// In addition to the main goroutine, launch two additional goroutine
//    - each additional goroutine should print something out
// Use waitgroups to ensure, goroutines finishes before program finishes

import (
	"fmt"
	"runtime"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s function in process\n", s)
	fmt.Println("Say CPU:", runtime.NumCPU())
	fmt.Println("Say GoRoutine:", runtime.NumGoroutine())
}

func listen(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s function in process\n", s)
	fmt.Println("Listen CPU:", runtime.NumCPU())
	fmt.Println("Listen GoRoutine:", runtime.NumGoroutine())
}

func main() {
	fmt.Println("Initial CPU:", runtime.NumCPU())
	fmt.Println("Initial GoRoutine:", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go say("say", &wg)
	go listen("listen", &wg)

	wg.Wait()
	fmt.Println("End CPU:", runtime.NumCPU())
	fmt.Println("End GoRoutine:", runtime.NumGoroutine())
}
