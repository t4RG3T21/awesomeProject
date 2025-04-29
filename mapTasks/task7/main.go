package main

//задача 7

import "fmt"

func removeDuplicates(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, str := range input {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	stringSlice := []string{"a", "b", "c", "d", "c", "c"}

	fmt.Println(removeDuplicates(stringSlice))
}
