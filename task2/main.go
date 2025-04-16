package main

import (
	"fmt"
)

//2. Сложение двух чисел
//Создайте программу, которая запрашивает у пользователя два числа и выводит их сумму.

func main() {
	fmt.Println(sumOf(2, 8))
}
func sumOf(x, y int) int {
	return x + y
}
