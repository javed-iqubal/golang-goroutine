package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

// Read
func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Write
func (c *Cache) Set(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {

	cache := &Cache{data: make(map[string]int)}

	// Simulate read-heavy workload (90% read)
	var wg sync.WaitGroup

	// 90 reader goroutine
	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cache.Get("key") // Fast- multiple readers
			}
		}(i)
	}

	// 10 write goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 10) // simulate slow write
			cache.Set("key", id)

		}(i)
	}

	start := time.Now()
	wg.Wait()
	fmt.Printf("RWMutex ::: %d Readers,  %d Writers took %v.\n", 90, 10, time.Since(start))

}
