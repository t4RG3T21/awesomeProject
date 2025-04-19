package main

//6. Поиск максимального числа
//Создайте программу, которая находит максимальное число в массиве целых чисел.
import "fmt"

func FindMax(numbers []int) int {

	max := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	return max
}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Println("Максимальное число:", FindMax(numbers))
}
