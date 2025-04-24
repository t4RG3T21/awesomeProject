package main

import "fmt"

func main() {
	strings := []string{"a", "b"}
	fmt.Println(strings)

	someFunc(&strings)
	fmt.Println(strings)
}

func someFunc(slice *[]string) {
	*slice = append(*slice, "modified")
}
