package main

import "fmt"

// if there is no publisher only subscriber then it is going to deadlock!
/**
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.subscriber(0x29b03f028738?)
        /home/javed/vscode-workspace/golang-goroutine/05goroutine/main.go:22 +0x25
main.main()
        /home/javed/vscode-workspace/golang-goroutine/05goroutine/main.go:30 +0x69
exit status 2
*/

func publisher(channel chan string) {
	//channel <- "Hello! How are you?"
}

func subscriber(channel chan string) {
	message := <-channel
	fmt.Println("Message from channel is: ", message)
}

func main() {
	channel := make(chan string)

	go publisher(channel)
	subscriber(channel)
}
