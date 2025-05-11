package main

import "fmt"

func inner() {
	defer fmt.Println("Inner defer")
	fmt.Println("Inner function")
}

func outer() {
	defer fmt.Println("Outer defer")
	inner()
	fmt.Println("Outer function")
}

func main() {
	outer()
}
