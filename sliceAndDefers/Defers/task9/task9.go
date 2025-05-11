package main

import "fmt"

func conditionalReturn(flag bool) (result int) {
	defer func() { result++ }()
	if flag {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(conditionalReturn(true))  // 2
	fmt.Println(conditionalReturn(false)) // 1
}
