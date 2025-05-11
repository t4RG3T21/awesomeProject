package main

import "fmt"

func GroupByLength(strs []string) map[int][]string {
	groups := make(map[int][]string)

	for _, s := range strs {
		length := len(s)
		groups[length] = append(groups[length], s)
	}

	return groups
}

func main() {
	example := []string{"cat", "dog", "bird", "go", "owl"}
	result := GroupByLength(example)

	for length, words := range result {
		fmt.Printf("%d букв: %v\n", length, words)
	}
}
