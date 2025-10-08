package main

import (
	"fmt"
	"time"
)

//6. Таймауты в каналах
//Используйте select с time.After, чтобы:
//Ждать данных из канала не более 2 секунд.
//Если данных нет — вывести "Timeout!".
//Цель: реализовать таймауты для операций с каналами.

func main() {
	// Создаем канал для данных
	dataChan := make(chan string)

	// Запускаем горутину, которая имитирует долгую операцию
	go func() {
		// Имитируем работу, которая занимает случайное время от 1 до 3 секунд
		duration := time.Duration(1+time.Now().UnixNano()%3) * time.Second
		time.Sleep(duration)
		dataChan <- "результат операции"
	}()

	// Используем select для ожидания с таймаутом
	select {
	case result := <-dataChan:
		// Если данные пришли до таймаута
		fmt.Println("Успешно получено: ", result)
	case <-time.After(2 * time.Second):
		// Если прошло 2 секунды и данные не пришли
		fmt.Println("Таймаут! Данные не получены за 2 секунды")
	}

	// Дополнительный пример: повторная попытка с несколькими таймаутами
	fmt.Println("\n=== Повторная попытка ===")

	// Создаем новый канал для второй попытки

	dataChan2 := make(chan string)

	go func() {
		// На этот раз операция занимает 3 секунды (больше таймаута)
		time.Sleep(3000 * time.Millisecond)
		dataChan2 <- "результат второй операции"
	}()

	// Ждем с таймаутом 2 секунды
	select {
	case result := <-dataChan2:
		fmt.Println("Успешно получено: ", result)

	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут при второй попытке!")
	}
}
