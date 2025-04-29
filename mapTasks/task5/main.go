package main

//задача 5

import "fmt"

func addGrade(db map[string]map[string]int, student, subject string, grade int) {
	if _, exists := db[student]; !exists {
		db[student] = make(map[string]int)
	}
	db[student][subject] = grade
}

func main() {
	db := make(map[string]map[string]int)

	addGrade(db, "Alice", "Math", 90)
	addGrade(db, "Alice", "Physics", 85)
	addGrade(db, "Bob", "History", 78)

	fmt.Println("База данных:")
	for student, subjects := range db {
		fmt.Printf("%s: %v\n", student, subjects)
	}
}
