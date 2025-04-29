package main

//задача 4

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("listen", "silent")) // true
	fmt.Println(isAnagram("hello", "bellow"))  // false
	fmt.Println(isAnagram("Apple", "apple"))   // true
	fmt.Println(isAnagram("test", "tes"))      // false
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	charCount := make(map[rune]int)

	for _, char := range s1 {
		charCount[char]++
	}

	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	return true
}
