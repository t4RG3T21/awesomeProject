package main

import "fmt"

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog)
	MakeSound(cat)
}

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Гав!"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Мяу!"
}

func MakeSound(a Animal) {
	fmt.Println(a.Sound())
}
