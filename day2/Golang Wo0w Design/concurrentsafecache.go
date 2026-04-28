
package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache is a thread-safe in-memory key-value store
type Cache struct {

	data map[string]interface{}
	mu   sync.RWMutex

}

// NewCache initializes the cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set adds or updates a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Delete removes a key from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Size returns the current number of keys
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func main() {
	cache := NewCache()

	// Start concurrent writers
	for i := 0; i < 5; i++ {

		go func(id int) {
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				cache.Set(key, fmt.Sprintf("value-%d", j))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// Start concurrent readers
	for i := 0; i < 3; i++ {
		
		go func(id int) {
			for {
				val, ok := cache.Get("key-1-5")
				if ok {
					fmt.Printf("[Reader %d] Found: %v\n", id, val)
				}
				time.Sleep(15 * time.Millisecond)
			}
		}(i)
	}

	// Let goroutines work for a bit
	time.Sleep(2 * time.Second)
	fmt.Printf("✅ Final Cache Size: %d\n", cache.Size())
}
