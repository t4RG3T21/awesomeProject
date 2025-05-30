package main

import "fmt"

type Calculator struct{}

func (c Calculator) Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	calc := Calculator{}

	fmt.Println(calc.Sum(1, 2, 3)) // 6
	fmt.Println(calc.Sum(10, 20))  // 30
}
