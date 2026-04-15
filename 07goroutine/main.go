package main

import "fmt"

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "!"

		// closing the channel so that receiver will not be in forever wait e.i deadlock
		close(channel)
	}()

	// reading multiple message from channel

	for value := range channel {
		fmt.Println(value)
	}

}
