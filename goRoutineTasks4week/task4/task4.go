package main

import (
	"context"
	"fmt"
	"sync"
)

//4. Построение конвейеров обработки данных
//Реализуйте конвейер (pipeline), где:
//Генератор создаёт поток данных (например, числа от 1 до 100)
//Фильтр пропускает только данные, удовлетворяющие условию (например, чётные числа)
//Обработчик преобразует данные (например, умножает на 10)
//Все этапы должны поддерживать отмену через context.Context
//При ошибке на любом этапе весь конвейер останавливается

// Генератор данных
func generator(ctx context.Context, out chan<- int) {
	defer close(out)
	for i := 1; i <= 100; i++ {
		select {
		case <-ctx.Done():
			return
		case out <- i:
		}
	}
}

// Фильтр данных
func filter(ctx context.Context, in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		select {
		case <-ctx.Done():
			return
		case out <- num:
			if num%2 != 0 {
				continue // Пропускаем нечетные числа
			}
		}
	}
}

// Обработчик данных
func handler(ctx context.Context, in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		select {
		case <-ctx.Done():
			return
		case out <- num * 10:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отключение контекста по завершению

	dataChan := make(chan int)
	filteredChan := make(chan int)
	processedChan := make(chan int)

	var wg sync.WaitGroup

	// Запускаем горутины
	wg.Add(3)
	go func() {
		defer wg.Done()
		generator(ctx, dataChan)
	}()
	go func() {
		defer wg.Done()
		filter(ctx, dataChan, filteredChan)
	}()
	go func() {
		defer wg.Done()
		handler(ctx, filteredChan, processedChan)
	}()

	// Чтение и вывод результатов
	go func() {
		for result := range processedChan {
			fmt.Println(result)
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}
