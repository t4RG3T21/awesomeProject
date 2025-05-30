package main

import "fmt"

func PrintType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type: int | Value: %d\n", v)
	case string:
		fmt.Printf("Type: string | Value: %s\n", v)
	case bool:
		fmt.Printf("Type: bool | Value: %t\n", v)
	case float64:
		fmt.Printf("Type: float64 | Value: %f\n", v)
	case rune:
		fmt.Printf("Type: rune | Value: %c\n", v)
	default:
		fmt.Printf("Type: %T | Value: %v\n", v, v)
	}
}

func main() {
	PrintType(42)       // int
	PrintType("gopher") // string
}
