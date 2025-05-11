package main

import "fmt"

func closureDefer() {
	x := 1
	defer func() { fmt.Println("Defer:", x) }() // Замыкание захватывает текущее x
	x = 2
	fmt.Println("Main:", x)
}

func main() {
	closureDefer()
}
