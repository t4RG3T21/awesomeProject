package main

//задача 1

import "fmt"

func countElements(words []string) map[string]int {
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	input := []string{"a", "b", "a", "c", "b"}
	result := countElements(input)
	fmt.Println(result) // map[a:2 b:2 c:1]
}
