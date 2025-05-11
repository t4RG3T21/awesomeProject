package main

//1. Подсчёт частоты элементов в слайсе
//Напишите функцию CountFrequency, которая принимает слайс строк и возвращает мапу,
//где ключи — элементы слайса, а значения — количество их вхождений.

import "fmt"

func CountFrequency(slice []string) map[string]int {
	frequency := make(map[string]int)

	for _, item := range slice {
		frequency[item]++
	}

	return frequency
}

func main() {
	input := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result := CountFrequency(input)
	fmt.Println(result) // map[apple:3 banana:2 orange:1]
}
