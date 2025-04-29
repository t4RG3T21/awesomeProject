package main

import "fmt"

func main() {
	q := map[string]int{"a": 1, "b": 2, "c": 3}
	w := map[string]int{"a": 3, "w": 2, "e": 3}

	m := mergeMap(q, w)

	fmt.Println(m)
}

func mergeMap(a, b map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range a {
		result[k] += v
	}
	for k, v := range b {
		result[k] += v
	}
	return result
}
