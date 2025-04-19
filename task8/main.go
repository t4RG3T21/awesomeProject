package main

//8. Сумма элементов массива
//Реализуйте программу, которая вычисляет сумму всех элементов в массиве.

import "fmt"

func SumArray(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(SumArray(arr))
}
