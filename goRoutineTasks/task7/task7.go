package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32     // Общий счётчик (используем int32 для atomic)
	var wg sync.WaitGroup // Наш механизм ожидания

	// Запускаем 4 горутины
	for i := 1; i <= 4; i++ {
		wg.Add(1) // Регистрируем новую горутину

		go func(workerID int) {
			defer wg.Done() // Гарантированно уменьшаем счётчик при завершении

			// Имитация работы
			time.Sleep(time.Duration(workerID*100) * time.Millisecond)

			// Безопасное увеличение счётчика
			newValue := atomic.AddInt32(&counter, 1)

			fmt.Printf("Рабочий %d: увеличил счётчик до %d\n", workerID, newValue)
		}(i)
	}

	fmt.Println("Менеджер: жду завершения всех рабочих...")
	wg.Wait() // Ожидаем завершения ВСЕХ горутин

	fmt.Printf("\nВсе рабочие завершили задание!\nИтоговое значение счётчика: %d\n", counter)
}
