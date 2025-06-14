package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int // Наша общая переменная
	var wg sync.WaitGroup
	var mu sync.Mutex // Создаём мьютекс для синхронизации

	numGoroutines := 5 // Количество горутин

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Имитация сложной работы
			time.Sleep(time.Duration(id*100) * time.Millisecond)

			mu.Lock() // Блокируем доступ для других горутин
			{
				// Критическая секция - работа с общей переменной
				temp := counter                  // Чтение
				temp++                           // Модификация
				time.Sleep(1 * time.Millisecond) // Имитация задержки
				counter = temp                   // Запись

				fmt.Printf("Горутина %d: counter = %d\n", id, counter)
			}
			mu.Unlock() // Разблокируем доступ
		}(i)
	}

	wg.Wait()
	fmt.Println("\nИтоговое значение counter:", counter)
}
