package main

import (
	"fmt"
)

// 1. Создаём структуру для ошибки деления
type DivError struct {
	dividend float64
	divisor  float64
}

// 2. Реализуем интерфейс error
func (e *DivError) Error() string {
	return fmt.Sprintf(
		"деление на 0: dividend=%.2f, divisor=%.2f",
		e.dividend,
		e.divisor,
	)
}

// 3. Функция деления с проверкой ошибок
func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// 4. Возвращаем кастомную ошибку
		return 0, &DivError{
			dividend: dividend,
			divisor:  divisor,
		}
	}
	return dividend / divisor, nil
}

func main() {
	// 5. Пример использования
	_, err := Divide(10, 0)
	if err != nil {
		fmt.Println(err) // деление на 0: dividend=10.00, divisor=0.00
	}
}
