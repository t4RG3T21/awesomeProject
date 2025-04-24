package main

//5я задача
import (
	"fmt"
)

func contains(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}

func getMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Содержит 'banana':", contains(fruits, "banana")) // true
	fmt.Println("Содержит 'grape':", contains(fruits, "grape"))   // false

	fmt.Println(getMax(3, 7, 2)) // 7
	fmt.Println(getMax())        // 0
}
