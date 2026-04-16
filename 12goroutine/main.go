package main

import (
	"fmt"
	"time"
)

func main() {

	// create two channel
	ch1 := make(chan int)
	ch2 := make(chan error)

	// first go routine send number
	go func() {
		time.Sleep(time.Millisecond * 2)
		ch1 <- 42
	}()

	// second go routine send error
	go func() {
		time.Sleep(time.Millisecond * 2)
		ch2 <- fmt.Errorf("Somthing bad happened!")
	}()

	fmt.Println("Wainting for result and error")

	// select statement which wait for int and error channel, who came first that result will be return.
	select {
	case value := <-ch1:
		fmt.Println("Result received: ", value)
	case err := <-ch2:
		fmt.Println("Error received: ", err)
	case <-time.After(time.Millisecond * 3):
		fmt.Println("Time Out.")
	}

}
