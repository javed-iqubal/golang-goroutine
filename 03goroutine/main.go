package main

import "fmt"

func main() {

	channel := make(chan string)
	senderMessage := "Welcome message in go channel"

	// sender logic
	go func() {
		channel <- senderMessage
		fmt.Println("Producer has published data...")
	}()

	// receiver logic
	receiverMessage := <-channel
	fmt.Println("Received message from channel is : ", receiverMessage)
}
