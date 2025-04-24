package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums) // [1 2 3 4 5 6 7 8 9 10]

	doubleValue(nums)
	fmt.Println(nums) // [2 4 6 8 10 12 14 16 18 20]
}

func doubleValue(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}
