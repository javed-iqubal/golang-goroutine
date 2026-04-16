package main

import (
	"fmt"
	"time"
)

func fetchFromCloud(serverName string, ch chan<- string, delay time.Duration) {

	time.Sleep(delay)
	ch <- "Data from " + serverName

}

func main() {

	serverA := make(chan string)
	serverB := make(chan string)

	go fetchFromCloud("http://www.proxy1.com", serverA, time.Millisecond*2)
	go fetchFromCloud("http://www.proxy2.com", serverB, time.Millisecond*2)

	// select statement acts as a race: first channel to send a value wins

	select {
	case messageOne := <-serverA:
		fmt.Println("SUCCESS : ", messageOne)
	case messageTwo := <-serverB:
		fmt.Println("SUCCESS: ", messageTwo)
	case <-time.After(time.Millisecond * 3): // safety to avoid deadlock
		fmt.Println("TimeOut: Both servers are too slow.")
	}

}
