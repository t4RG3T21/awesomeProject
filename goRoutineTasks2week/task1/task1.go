package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, exists := c.items[key]
	return val, exists
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *Cache) Update(key string, updater func(string) string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if old, exists := c.items[key]; exists {
		c.items[key] = updater(old)
	}
}

func main() {
	cache := NewCache()

	// Запись данных
	cache.Set("name", "Alice")
	cache.Set("role", "admin")

	// Параллельное чтение (10 читателей)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if val, ok := cache.Get("name"); ok {
				fmt.Printf("Reader %d: %s\n", id, val)
			}
		}(i)
	}

	// Обновление в отдельной горутине
	go func() {
		time.Sleep(100 * time.Millisecond)
		cache.Update("name", func(s string) string {
			return "Bob"
		})
		fmt.Println("Updated name to Bob")
	}()

	wg.Wait()
	fmt.Println("All readers completed")
}
