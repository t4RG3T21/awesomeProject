package main

import "fmt"

func main() {
	// 1. создал срез
	s := make([]int, 3, 5)
	fmt.Printf("Исходный: len=%d, cap=%d\n", len(s), cap(s))

	// 2. добавил элемент
	s = append(s, 4)
	fmt.Printf("После добавления 4: len=%d, cap=%d\n", len(s), cap(s))

	// 3. добавил два элемента за один вызов append
	s = append(s, 5, 6)
	fmt.Printf("После добавления 6: len=%d, cap=%d\n", len(s), cap(s))
}
