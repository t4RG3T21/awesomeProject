package main

//7. Проверка палиндрома
//Напишите функцию, которая проверяет, является ли строка палиндромом.
import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("А роза упала на лапу Азора")) // true
	fmt.Println(IsPalindrome("Hello"))                      // false
}
