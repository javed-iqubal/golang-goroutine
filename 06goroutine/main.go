package main

import "fmt"

func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "!"
		//close(channel)
	}()

	// reading multiple message from channel

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel)
}
