package main

import (
	"fmt"
	"time"
)

func Hello(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {

	fmt.Println("main go routine started")

	// launch two go routine it run independently..
	go Hello("Hello")
	go Hello("Welcome")

	// block or wait the main go routine, otherwise main exits before Hello go routine print
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("main go routine ended")
}
