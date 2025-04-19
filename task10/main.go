package main

//10. Уникальные элементы
//Создайте программу, которая принимает массив и возвращает новый массив,
//содержащий только уникальные элементы.

import "fmt"

func UniqueElements(arr []int) []int {
	result := []int{}

	// проверка, был ли элемент уже добавлен
	for i := 0; i < len(arr); i++ {
		isUnique := true
		for j := 0; j < len(result); j++ {
			if arr[i] == result[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			result = append(result, arr[i])
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(UniqueElements(arr)) // [1 2 3 4 5]
}
