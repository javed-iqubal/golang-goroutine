package main

import (
	"fmt"
	goroutine "goroutine/routinepkg"
	"sync"
)

func main() {

	fmt.Println("Main goroutine running...")

	//A WaitGroup is a counting semaphore typically used to wait for a group of
	//goroutines or tasks to finish.
	var wg sync.WaitGroup

	counter := 10

	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go goroutine.HelloNew("hello", &wg)
	}
	fmt.Println("After loop::: main goroutine running...")
	//pause the main or any goroutine to finish
	wg.Wait()
	fmt.Println("Main goroutine finished")
}
