package main

import (
	"fmt"
)

// Этап 1: Генератор числа - производит числа от 1 до n
// Возвращает канал только для чтения, из которого можно получать числа
func generator(n int) <-chan int {
	out := make(chan int) // создаем канал для отправки чисел

	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// Этап 2: Умножение чисел - принимает числа из канала, умножает их на 2 и отправляет дальше
// Принимает канал только для чтения, возвращает канал только для чтения
func multiplier(in <-chan int) <-chan int {
	out := make(chan int) // создаем канал для отправки результатов

	go func() {
		for num := range in { // читаем числа из входного канала
			result := num * 2
			out <- result
		}
		close(out)
	}()
	return out
}

// Этап 3: Вывод результатов - принимает числа из канала и выводит их
// Принимает канал только для чтения
func printer(in <-chan int) {
	for result := range in {
		fmt.Printf("Результат: %d\n", result)
	}
	fmt.Println("Все результаты обработаны")
}

// дополнительный этап: фильтр - пропускает только четные числа
func evenFilter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("=== Простой конвейер: Генератор -> Умножитель -> Принтер ===")

	// Создаем конвейер из трех этапов:
	// 1. Генератор создает числа от 1 до 5
	// 2. умножитель умножает каждое число на 2
	// 3. принтер выводит результаты
	numbers := generator(5)
	multiplied := multiplier(numbers)
	printer(multiplied)

	fmt.Println("\n=== Расширенный конвейер с фильтром ===")

	// более сложный конвейер с дополнительным этапом фильтрации:
	// генератор -> фильтр (только четные) -> умножитель -> принтер
	numbers2 := generator(10)              // числа от 1 до 10
	evenNumbers := evenFilter(numbers2)    // фильтруем четные
	multiplied2 := multiplier(evenNumbers) // умножаем четные числа на 2
	printer(multiplied2)                   // выводим

	fmt.Println("\n=== Конвейер в одной строке ===")
	// конвейер можно собрать и в одной строке:
	printer(multiplier(generator(3)))
}
