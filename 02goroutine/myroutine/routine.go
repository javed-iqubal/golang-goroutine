package goroutine

import (
	"fmt"
	"sync"
)

func HelloNew(message string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
	}
	defer wg.Done()
}
