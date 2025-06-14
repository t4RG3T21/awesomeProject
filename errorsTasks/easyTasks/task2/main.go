package main

import (
	"fmt"
	"math"
)

type NegativeNumberError struct {
	Number float64
}

func (e *NegativeNumberError) Error() string {
	return fmt.Sprintf("невозможно вычислить квадратный корень из отрицательного числа: %.2f", e.Number)
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, &NegativeNumberError{Number: n}
	}
	return math.Sqrt(n), nil
}

func main() {
	numbers := []float64{4, -9, 0, 25, -1}

	for _, num := range numbers {
		result, err := Sqrt(num)
		if err != nil {
			if e, ok := err.(*NegativeNumberError); ok {
				fmt.Printf("Ошибка для %.2f: %s\n", num, e.Error())
			} else {
				fmt.Println("Неизвестная ошибка:", err)
			}
		} else {
			fmt.Printf("√%.2f = %.2f\n", num, result)
		}
	}
}
