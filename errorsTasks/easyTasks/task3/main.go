package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testCases := []string{
		"123",                            // валидное число
		"12.3",                           // синтаксическая ошибка
		"999999999999999999999999999999", // выход за пределы диапазона
		"abc",                            // синтаксическая ошибка
		"",                               // синтаксическая ошибка
	}

	for _, str := range testCases {
		fmt.Printf("Пытаемся преобразовать: '%s'\n", str)

		num, err := strconv.Atoi(str)

		if err != nil {
			switch {
			case strings.Contains(err.Error(), "syntax"):
				// ошибка парсинга (синтаксическая)
				fmt.Printf("  Ошибка парсинга: '%s' не является целым числом\n", str)

			case strings.Contains(err.Error(), "range"):
				// ошибка диапазона
				fmt.Printf("  Ошибка диапазона: '%s' выходит за пределы допустимого диапазона\n", str)

			default:
				// неизвестный тип ошибки
				fmt.Printf("  Неизвестная ошибка: %v\n", err)
			}
		} else {
			// успешное преобразование
			fmt.Printf("  Успех: преобразовано в %d\n", num)
		}

		fmt.Println("---")
	}
}
