package main

//9. Числа Фибоначчи
//Напишите функцию, которая генерирует первые N чисел Фибоначчи.

import "fmt"

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fib := make([]int, n)

	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	fmt.Println(Fibonacci(10)) // [0 1 1 2 3 5 8 13 21 34]
}
