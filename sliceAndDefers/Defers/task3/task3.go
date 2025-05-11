package main

import "fmt"

func magicNumber() (result int) {
	defer func() { result = 42 }()

	result = 10
	return
}

func main() {
	fmt.Println(magicNumber())
}
