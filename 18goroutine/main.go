package main

import (
	"fmt"
)

// produces series of values - outbond channel
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

func sq(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // squre and write onto channel
		}
		close(out)
	}()

	return out

}

func main() {

	// var wg sync.WaitGroup

	// Fan out pattern: same channel input is distributed to multiple channels

	fmt.Println("Main started")
	in := gen(2, 3, 4, 5)

	// FAN-OUT 2 worker

	ch1 := sq(in)
	ch2 := sq(in)

	// consume each separatly

	for n1 := range ch1 {
		fmt.Println("ch1 -> ", n1)
	}
	for n2 := range ch2 {
		fmt.Println("ch2 => ", n2)
	}

	// wg.Wait()
	fmt.Println("Main End")
}
