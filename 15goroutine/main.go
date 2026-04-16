package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&counter, 1) // atomic operation
	fmt.Printf("Incrementing :: %d \n", counter)

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Value:: %d,   Finished! \n", counter)

}
