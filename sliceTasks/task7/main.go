package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(test1)

	deleteEvenNums(&test1)
	fmt.Println(test1)
}

func deleteEvenNums(slice *[]int) {
	writeIndex := 0
	for _, value := range *slice {
		if value%2 != 0 {
			(*slice)[writeIndex] = value
			writeIndex++
		}
	}
	*slice = (*slice)[:writeIndex]
}
