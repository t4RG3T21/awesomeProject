package main

//8. Подсчёт среднего значения в слайсе
//Напишите функцию Average, которая принимает слайс чисел и возвращает их среднее значение. Если слайс пуст, возвращает 0.

import "fmt"

func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += float64(num) // Преобразуем int в float64 для точного деления
	}
	return sum / float64(len(numbers)) // Возвращаем среднее как float64
}

func main() {
	// Примеры использования
	fmt.Println(Average([]int{1, 2, 3, 4, 5})) // 3.0
	fmt.Println(Average([]int{2, 4}))          // 3.0
	fmt.Println(Average([]int{-1, 0, 1}))      // 0.0
	fmt.Println(Average([]int{}))              // 0
}
