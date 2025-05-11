package main

import "fmt"

func multipleDefers() {
	// Добавляем три отложенных вызова
	defer fmt.Println("Defer 1 - First In")  // 3-й по порядку выполнения
	defer fmt.Println("Defer 2 - Second In") // 2-й по порядку выполнения
	defer fmt.Println("Defer 3 - Last In")   // 1-й по порядку выполнения

	fmt.Println("Start main logic")
}

func main() {
	fmt.Println("Start program")
	multipleDefers()
	fmt.Println("End program")
}
