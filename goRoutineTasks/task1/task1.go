package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello, Go!")
	}()

	time.Sleep(300 * time.Millisecond)

	fmt.Println("Главная горутина завершена")
}
