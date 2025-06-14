package main

import (
	"fmt"
	"sync"
	"time"
)

// SquareCalculator структура для безопасного вычисления квадрата
type SquareCalculator struct {
	result int           // Разделяемая переменная для результата
	mu     sync.Mutex    // Мьютекс для синхронизации доступа
	done   chan struct{} // Канал для сигнала о завершении
}

// NewSquareCalculator создает новый экземпляр калькулятора
func NewSquareCalculator() *SquareCalculator {
	return &SquareCalculator{
		done: make(chan struct{}),
	}
}

// ComputeSquare запускает горутину для вычисления квадрата числа
func (sc *SquareCalculator) ComputeSquare(n int) {
	go func() {
		// Имитация долгих вычислений
		time.Sleep(500 * time.Millisecond)

		sc.mu.Lock()      // Блокируем доступ к разделяемой переменной
		sc.result = n * n // Вычисляем и сохраняем результат
		sc.mu.Unlock()    // Разблокируем доступ

		close(sc.done) // Сигнализируем о завершении
	}()
}

// GetResult возвращает результат вычислений, дожидаясь их завершения
func (sc *SquareCalculator) GetResult() int {
	<-sc.done            // Ждем сигнала о завершении
	sc.mu.Lock()         // Блокируем для безопасного чтения
	defer sc.mu.Unlock() // Гарантируем разблокировку
	return sc.result
}

func main() {
	calculator := NewSquareCalculator()

	fmt.Println("Запускаем вычисление квадрата для числа 7...")
	calculator.ComputeSquare(7)

	// Пока горутина работает, можем выполнять другие задачи
	fmt.Println("Горутина работает... можно делать другие операции...")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Выполняем другие задачи...")

	// Получаем результат (будет ждать завершения горутины)
	result := calculator.GetResult()
	fmt.Printf("Результат: 7² = %d\n", result)

	// Демонстрация повторного использования
	fmt.Println("\nЗапускаем вычисление для числа 12...")
	calculator.ComputeSquare(12)
	fmt.Println("Ждем результат...")
	fmt.Printf("Результат: 12² = %d\n", calculator.GetResult())
}
