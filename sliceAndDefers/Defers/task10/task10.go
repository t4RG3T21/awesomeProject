package main

import (
	"fmt"
	"time"
)

func longFunction() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()
	time.Sleep(2 * time.Second)
}

func main() {
	longFunction()
}
