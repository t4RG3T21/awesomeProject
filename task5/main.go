package main

import "fmt"

//5. Обратная строка
//Напишите программу, которая принимает строку от пользователя
//и выводит ее в обратном порядке.

func main() {
	fmt.Print(ReverseString("string"))
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
