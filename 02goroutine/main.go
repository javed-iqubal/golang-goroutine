package main

import (
	"fmt"
	goroutine "mygoroutine/myroutine"
	"sync"
)

func main() {

	fmt.Println("Start func main")

	/**
	A WaitGroup is a counting semaphore typically used to wait for a group of goroutine
	or task to finish.
	*/
	var wg sync.WaitGroup

	// How many goroutines you are going to be launch
	wg.Add(2)
	go goroutine.HelloNew("GoLang", &wg)
	go goroutine.HelloNew("Python", &wg)
	fmt.Println("func main running....")

	// pause the main or any goroutine finish
	wg.Wait()

	fmt.Println("End func main")
}
