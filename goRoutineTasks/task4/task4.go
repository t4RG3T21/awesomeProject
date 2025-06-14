package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) { // Параметр n — локальная копия
			defer wg.Done()
			fmt.Printf("Правильно (параметр): %d\n", n)
		}(i) // Передаём i В МОМЕНТ итерации
	}

	wg.Wait()
}
