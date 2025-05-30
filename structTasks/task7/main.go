package main

import (
	"fmt"
	"sort"
)

// 1. Объявляем структуру Student

type Student struct {
	Name  string
	Grade int
}

// 2. Создаем тип-обертку для реализации sort.Interface

type ByGrade []Student

// 3. Реализуем методы интерфейса sort.Interface

func (s ByGrade) Len() int {
	return len(s)
}

func (s ByGrade) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByGrade) Less(i, j int) bool {
	return s[i].Grade < s[j].Grade // Сортировка по возрастанию оценок
}

func main() {
	// 4. Создаем и инициализируем слайс студентов
	students := []Student{
		{"Alice", 90},
		{"Bob", 75},
		{"Charlie", 85},
		{"David", 82},
	}

	fmt.Println("До сортировки:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Grade)
	}

	// 5. Сортируем слайс
	sort.Sort(ByGrade(students))

	fmt.Println("\nПосле сортировки:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Grade)
	}
}
