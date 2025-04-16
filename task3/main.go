package main

import "fmt"

//3. Число четное или нечетное
//Напишите программу, которая проверяет,
//является ли введенное пользователем число четным или нечетным.

func main() {
	fmt.Println(factorial(5))
}
func factorial(x uint) uint {
	var result uint = 1
	var i uint
	for i = 1; x >= i; i++ {
		result = i * result
	}
	return result
}
