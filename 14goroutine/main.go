package main

import (
	"fmt"
	"sync"
)

// object which hold data
type SafeCounter struct {
	mu    sync.Mutex
	count int // data to be protected from multiple go routine access
}

// function to return counter

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

// function to increment count
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

func main() {

	// create SafeCounter
	counter := SafeCounter{
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()

	}
	wg.Wait()
	fmt.Printf("Counter value: %d \n", counter.Value())

}
