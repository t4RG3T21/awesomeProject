package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	reverseSlice(slice1)
	fmt.Println(slice1)
}

func reverseSlice(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}
}
