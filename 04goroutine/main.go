package main

import "fmt"

func sayHello(channel chan string) {
	channel <- "Hello! How are you?"
}

func receiver(channel chan string) {
	message := <-channel
	fmt.Println("Message from channel is: ", message)
}

func main() {
	channel := make(chan string)

	go sayHello(channel)
	receiver(channel)
}
