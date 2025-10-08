package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//1. Таймауты и отмена операций
//Напишите функцию, которая:
//Выполняет долгую операцию (например, time.Sleep или цикл вычислений)
//Принимает context.Context для контроля выполнения
//Прекращает работу, если контекст отменён (например, по таймауту 2 секунды)
//Возвращает ошибку, если операция не завершилась вовремя

// longOperation выполняет долгую операцию с поддержкой отмены через контекст
func longOperation(ctx context.Context) error {
	fmt.Println("Начало долгой операции")

	// имитируем долгую операцию (например, запрос к API или вычисления)
	for i := 0; i < 10; i++ {
		// проверяем, не отменен ли контекст
		select {
		case <-ctx.Done():
			// контекст отменен, прекращаем работу
			fmt.Println("Операция прервана по таймауту или отмене")
			return ctx.Err()
		default:
			// контекст активен, продолжаем работу
			fmt.Printf("Выполняется шаг %d\n", i+1)
			time.Sleep(500 * time.Millisecond) // имитация работы
		}
	}

	fmt.Println("Долгая операция успешно завершена")
	return nil
}

func main() {
	// вариант 1: контекст с таймаутом 2 секунды
	fmt.Println("=== Тест с таймаутом 2 секунды ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1() // всегда откладываем отмену для освобождения ресурсов

	if err := longOperation(ctx1); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Ошибка: превышено время ожидания")
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}

	// вариант 2: контекст с ручной отменой
	fmt.Println("\n=== Тест с ручной отменой ===")
	ctx2, cancel2 := context.WithCancel(context.Background())

	// запускаем операцию в горутине
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("отменяем операцию...")
		cancel2()
	}()

	if err := longOperation(ctx2); err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("Ошибка: операция отменена вручную")
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}
}
