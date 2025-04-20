package main

//2я задача 2я часть

import (
	"fmt"
	"math"
)

func main() {
	L := 35.0

	// R = L / (2 * π)
	RValue := L / (2 * math.Pi)

	R := &RValue

	// S = π * R²
	S := math.Pi * (*R) * (*R)

	roundedS := math.Round(S*100) / 100
	roundedR := math.Round(*R*100) / 100

	fmt.Printf("Радиус: %.2f\n", roundedR) //форматирование, 2 знака после запятой
	fmt.Printf("Площадь: %.2f\n", roundedS)
}
