package main

import "fmt"

func main() {
	original := make([]int, 3, 4)
	original[0], original[1], original[2] = 10, 20, 30

	modified := addElByIndex(original, 1, 99)

	fmt.Printf(
		"Результат: %v\nДлина: %d\nЁмкость: %d",
		modified,
		len(modified),
		cap(modified),
	)
}

func addElByIndex(slice []int, index, value int) []int {
	if index < 0 || index > len(slice) {
		panic("Некорректный индекс")
	}

	slice = slice[:len(slice)+1]

	copy(slice[index+1:], slice[index:])

	slice[index] = value

	return slice
}
