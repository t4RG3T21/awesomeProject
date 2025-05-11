package main

import "fmt"

//7. Поиск уникальных элементов в слайсе
//Напишите функцию FindUnique, которая принимает слайс целых чисел и возвращает новый слайс,
//содержащий только уникальные элементы (встречающиеся 1 раз).

func main() {
	slice := []int{1, 2, 3, 3, 4, 4, 5, 6}
	fmt.Println(findUnique(slice))
}

func findUnique(nums []int) []int {
	counts := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums {
		counts[num]++
	}

	for _, num := range nums {
		if counts[num] == 1 {
			result = append(result, num)
		}
	}
	return result
}
