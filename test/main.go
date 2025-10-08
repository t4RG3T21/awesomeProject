package main

import (
	"fmt"
	"sync"
)

func main() {
	printN(10)
}

func printN(n int) {
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
