package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("Defer:", x) // Аргумент x оценивается как 1
	x = 2
	fmt.Println("Main:", x)
}
