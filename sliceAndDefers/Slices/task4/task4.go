package main

//4. Удаление дубликатов из слайса
//Напишите функцию RemoveDuplicates, которая принимает слайс строк и возвращает новый слайс без дубликатов.

import "fmt"

func main() {
	slice := []string{"a", "a", "b", "c", "d", "d", "e", "f", "f", "g"}

	test := removeDuplicates(slice)

	fmt.Println(test)
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
