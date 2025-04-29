package main

//задача 3

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 25, "c": 15, "d": 50, "e": 16}
	r := filterMap(m, 15)
	fmt.Println(r)
}

func filterMap(nums map[string]int, n int) map[string]int {
	result := make(map[string]int)
	for k, v := range nums {
		if v > n {
			result[k] += v
		}

	}
	return result
}
