package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s, Автор: %s", b.Title, b.Author)
}

func main() {
	book := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
	}

	fmt.Println(book) // Книга: Война и мир, Автор: Лев Толстой

	sciFi := Book{"Дюна", "Фрэнк Герберт"}
	fmt.Println(sciFi) // Книга: Дюна, Автор: Фрэнк Герберт
}
