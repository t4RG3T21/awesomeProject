package main

//задача 6

import "fmt"

func MapToSlices(m map[int]string) ([]int, []string) {
	keys := make([]int, 0, len(m))
	values := make([]string, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

func main() {
	input := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	keys, values := MapToSlices(input)
	fmt.Println("Ключи:", keys)
	fmt.Println("Значения:", values)
}
