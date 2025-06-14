package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			time.Sleep(time.Duration(n) * 100 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", n)
		}(i)
	}

	fmt.Println("Шеф ждёт...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
