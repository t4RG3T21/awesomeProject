package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.value.Add(1)
}

func (c *AtomicCounter) Value() int64 {
	return c.value.Load()
}

func main() {
	counter := AtomicCounter{}
	var wg sync.WaitGroup

	// Запускаем 1000 горутин для инкремента
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", counter.Value()) // Всегда 1000
}
