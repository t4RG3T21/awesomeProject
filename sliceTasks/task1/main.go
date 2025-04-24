package main

import "fmt"

func main() {
	createAndExtendSlice(7)
}

func createAndExtendSlice(n int) {
	original := make([]int, n)
	for i := 0; i < n; i++ {
		original[i] = i + 1
	}

	additional := make([]int, n)
	for i := 0; i < n; i++ {
		additional[i] = i + 1 + n
	}
	modified := append(original, additional...)
	fmt.Println(original) // original slice
	fmt.Println(modified) // modified slice
}
