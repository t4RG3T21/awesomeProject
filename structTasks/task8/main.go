package main

import (
	"fmt"
	"strings"
)

type Database interface {
	Save(data string)
	Read() string
}

type MockDB struct {
	storage []string // Слайс для хранения данных
}

func (db *MockDB) Save(data string) {
	db.storage = append(db.storage, data)
}

func (db *MockDB) Read() string {
	return strings.Join(db.storage, ", ")
}

func main() {
	var db Database = &MockDB{}

	db.Save("Hello")
	db.Save("Gopher")

	fmt.Println("Данные в БД:", db.Read())
}
