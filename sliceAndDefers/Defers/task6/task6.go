package main

import "fmt"

func namedReturn() (result int) {
	defer func() { result = 2 }() // Меняет result после return
	result = 1
	return
}

func main() {
	fmt.Println(namedReturn()) // 2
}
