package main

//5. Напишите функцию InvertMap, которая принимает мапу map[string]int и возвращает новую мапу map[int]string,
//где ключи и значения поменялись местами. Если в исходной мапе есть дублирующиеся значения, вернуть ошибку.

import (
	"errors"
	"fmt"
)

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	result1, err1 := InvertMap(m1)
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println("Инвертированная мапа:", result1)
		// Вывод: map[1:a 2:b 3:c]
	}
}

func InvertMap(original map[string]int) (map[int]string, error) {
	inverted := make(map[int]string)
	seenValues := make(map[int]bool)

	for key, value := range original {
		if seenValues[value] {
			return nil, errors.New("дублирующиеся значения в исходной мапе")
		}
		seenValues[value] = true
		inverted[value] = key
	}
	return inverted, nil
}
