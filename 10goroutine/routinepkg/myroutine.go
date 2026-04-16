package goroutine

import (
	"fmt"
	"sync"
)

func HelloNew(message string, wg *sync.WaitGroup) {

	fmt.Println(message)

	defer wg.Done()
}
