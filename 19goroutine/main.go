package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)

	}()

	return out
}

func squre(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out

}

// FAN-IN: merge channel

func merge(cs ...<-chan int) <-chan int {

	var wg sync.WaitGroup

	out := make(chan int)

	wg.Add(len(cs))

	for _, ch := range cs {

		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)

	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	in := gen(2, 3, 4, 5)

	//FAN-OUT
	ch1 := squre(in)
	ch2 := squre(in)

	// FAN_IN
	out := merge(ch1, ch2)

	// single consumer
	for c := range out {
		fmt.Println("megred: ", c)
	}

}
