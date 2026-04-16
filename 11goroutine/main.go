package main

import (
	"fmt"
	"sync"
	"time"
)

func processImage(image int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Image-%d started resizing...\n", image)
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("Image-%d resizing finished.\n", image)
}

func main() {

	fmt.Println("Main: Started....")
	var wg sync.WaitGroup

	imageToProcess := 5

	for i := 1; i <= imageToProcess; i++ {
		wg.Add(1)
		go processImage(i, &wg)
	}

	fmt.Println("Main: Waiting for all images to finish")
	wg.Wait()
	fmt.Println("Main: All images are processed. Ready to upload!")
}
