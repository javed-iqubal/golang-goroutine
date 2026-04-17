package main

import "fmt"

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

	// setup the pipeline
	c := gen(2, 3, 4)
	out := sq(c)

	for val := range out {
		fmt.Println("Value: ", val)
	}

	// short cut

	for value := range sq(gen(2, 3, 4, 5, 6, 7, 8, 9)) {
		fmt.Println("Value ------> ", value)
	}

}
