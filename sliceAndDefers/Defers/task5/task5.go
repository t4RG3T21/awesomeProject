package main

import "fmt"

func deferInLoop() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("Defer:", i) // Добавляется в стек, но выполняется после цикла
	}
	fmt.Println("End of loop")
}

func main() {
	deferInLoop()
}
