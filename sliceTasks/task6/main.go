package main

import "fmt"

func main() {
	test1 := []int{1, 2}
	test2 := []int{3, 4}

	fmt.Println(test1, test2)

	fmt.Println(concatSlices(test1, test2))
}

func concatSlices(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	result := make([]int, lenA+lenB)
	copy(result[:lenA], a)
	copy(result[lenA:], b)
	return result
}
