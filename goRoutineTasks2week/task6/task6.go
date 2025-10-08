package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type IDGenerator struct {
	counter atomic.Int64
}

func (gen *IDGenerator) NextID() int64 {
	return gen.counter.Add(1)
}

func main() {
	var generator IDGenerator
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				id := generator.NextID()
				fmt.Printf("Goroutine ID: %d\n", id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("All IDs generated")
}
