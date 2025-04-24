package main

import "fmt"

// 3. Удаление элемента по индексу
//Напишите функцию, которая удаляет элемент из среза по заданному индексу in-place (без возврата нового среза).
//Пример:
//Вход: slice = []int{10, 20, 30, 40}, index = 1
//Выход: [10 30 40]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	deleteByIndex(&nums, 4)
	fmt.Println(nums)
}

func deleteByIndex(slice *[]int, index int) {
	// представил, что индекс идет от 1, поэтому при передаче условно 4, пропадет конкретно 4й элемент
	*slice = append((*slice)[:index-1], (*slice)[index:]...)

}
