package main

import (
	"fmt"
)

// 3. Число четное или нечетное
// Напишите программу, которая проверяет, является ли введенное пользователем число четным или нечетным.
func main() {
	fmt.Println(isEven(4))
	fmt.Println(isEven(5))
}

func isEven(n int) bool {
	return n%2 == 0
}
