package main

//3. Объединение двух мап
//Напишите функцию MergeMaps,
//которая принимает две мапы (map[string]int) и возвращает новую мапу, объединяющую их.
//Если ключ присутствует в обеих мапах, выбирается значение из второй.

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 5, "c": 1, "d": 1}

	merged := mergeMaps(m1, m2)

	fmt.Println(merged)
}

func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int, len(m1)+len(m2))
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}
