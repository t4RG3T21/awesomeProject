package main

import (
	"context"
	"fmt"
	"time"
)

//1. Таймауты и отмена операций
//Напишите функцию, которая:
//Выполняет долгую операцию (например, time.Sleep или цикл вычислений)
//Принимает context.Context для контроля выполнения
//Прекращает работу, если контекст отменён (например, по таймауту 2 секунды)
//Возвращает ошибку, если операция не завершилась вовремя

// longOperation выполняет долгую операцию с возможностью отмены через контекст
func longOperation(ctx context.Context) error {
	// Имитируем долгую операцию (например, вычисления или запрос к БД)
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			// Контекст отменен (таймаут или явная отмена)
			return ctx.Err()
		default:
			// Продолжаем работу
			fmt.Printf("Выполнение шага %d...\n", i)
			time.Sleep(500 * time.Millisecond) // Имитация работы
		}
	}

	fmt.Println("Операция успешно завершена")
	return nil
}

func main() {
	// Создаем контекст с таймаутом 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Важно вызывать cancel для освобождения ресурсов

	// Запускаем операцию
	err := longOperation(ctx)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Программа завершена")
}
