package main

//2. Фильтрация слайса по условию
//Напишите функцию FilterEven, которая принимает слайс целых чисел и возвращает новый слайс, содержащий только чётные числа.

import "fmt"

func main() {
	testS := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(filterEven(testS))

}

func filterEven(slice []int) []int {
	result := make([]int, 0)
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}
