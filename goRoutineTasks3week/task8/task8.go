package main

import (
	"fmt"
	"time"
)

// producer создает числа и отправляетих в канал только для записи
// Параметр out имеет тип chan<- int, что означает "канал только для отправки"
// Функция может только отправлять данные в этот канал, но не читать из него
func producer(out chan<- int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Отправляем число %d\n", i)
		out <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(out)
}

// processor принимает числа из канала только для чтения
// обрабатывает их и отправляет результаты в канал только для записи
// Параметр in имеет тип <-chan int - "канал только для чтения"
// Параметр out имеет тип chan<- int - "канал только для записи"
func processor(in <-chan int, out chan<- int) {
	for num := range in {
		result := num * 2
		fmt.Printf("Обрабатываем %d -> %d\n", num, result)
		out <- result
	}
	close(out)
}

// consumer только читает данные из канала и выводит их
// параметр in имеет тип <-chan int - "канал только для чтения"
func consumer(in <-chan int) {
	for result := range in {
		fmt.Printf("Получен результат: %d\n", result)
	}
	fmt.Println("Все результаты получены")
}

func main() {
	numbers := make(chan int, 5) // канал для чисел
	results := make(chan int, 5) // канал для результатов

	// запускаем конвейер обработки

	// Producer отправляет числа в канал numbers
	go producer(numbers, 5) // отправляем 5 чисел

	// Processor обрабатывает числа и отправляет результаты
	go processor(numbers, results)

	// Consumer читает и выводит результаты
	consumer(results)

	fmt.Println("Программа завершена")
}
