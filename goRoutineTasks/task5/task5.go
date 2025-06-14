package main

import (
	"fmt"
	"sync"
)

func main() {
	var balance int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			balance++
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый баланс:", balance)
}
