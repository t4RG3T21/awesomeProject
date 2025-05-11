package main

import "fmt"

func deferExample() {
	fmt.Println("Start")       // 1. Выполняется первым
	defer fmt.Println("Defer") // 3. Выполняется ПОСЛЕ завершения функции
	fmt.Println("End")         // 2. Выполняется вторым
}

func main() {
	deferExample()
}
