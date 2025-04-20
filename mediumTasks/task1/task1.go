package main

//1я задача

import (
	"fmt"
	"strconv"
)

func main() {
	// Объявляем переменные
	strValue := "104"
	intValue := 35

	// Преобразуем строку в число
	numFromStr, _ := strconv.Atoi(strValue) // Игнорируем ошибку

	// Преобразуем число в строку
	strFromInt := fmt.Sprint(intValue)

	fmt.Println("Строка -> число:", numFromStr)
	fmt.Println("Число -> строка:", strFromInt)
}
