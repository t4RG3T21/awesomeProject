package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

//2. Параллельная обработка со сбором ошибок
//Реализуйте функцию, которая:
//Принимает слайс данных (чисел, строк и т.п.)
//Обрабатывает каждый элемент параллельно в горутинах
//Использует errgroup для управления горутинами
//Возвращает:
//Результаты обработки (если все элементы обработаны успешно)
//Первую возникшую ошибку (если хотя бы одна операция завершилась с ошибкой)

// processTime обрабатывает один элемент данных
func processItem(ctx context.Context, item string) (string, error) {
	// имитируем обработку элемента
	fmt.Printf("Обрабатывается элемент: %s\n", item)

	// преобразуем строку в число
	num, err := strconv.Atoi(item)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования '%s' в число: %w", item, err)
	}

	// имитируем возможную ошибку для определенного значения
	if num == 7 {
		return "", errors.New("элемент 7 приносит неудачу!")
	}

	// имитируем долгую работу
	select {
	case <-time.After(100 * time.Millisecond):
		// продолжаем обработку
	case <-ctx.Done():
		// контекст отменен, прекращаем работу
		return "", ctx.Err()
	}

	// возвращаем обработанный результат
	result := fmt.Sprintf("Обработано: %d", num*2)
	return result, nil
}

// parallelProcess обрабатывает слайс данных параллельно
func parallelProcess(ctx context.Context, items []string) ([]string, error) {
	// создаем errgroup с контекстом
	// g будет управлять группой горутин и автоматически отменять их при ошибке
	g, ctx := errgroup.WithContext(ctx)

	// слайс для хранения результатов
	results := make([]string, len(items))

	// обрабатываем каждый элемент в отдельной горутине
	for i, item := range items {
		i, item := i, item // захватываем текущее значение для горутины

		g.Go(func() error {
			// вызываем обработку элемента
			result, err := processItem(ctx, item)
			if err != nil {
				return err // если ошибка - возвращаем её
			}

			// сохраняем результат в соответствующую позицию
			results[i] = result
			return nil
		})
	}

	// ожидаем завершение всех горутин
	// если какая-то горутина вернула ошибку, Wait() вернет первую возникшую ошибку
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// если всё успешно - возвращаем результаты
	return results, nil
}

func main() {
	// Тестовые данные
	items := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	// создаем контекст с таймаутом на всякий случай
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("=== Начало параллельной обработки ===")

	// вызываем параллельную обработку
	results, err := parallelProcess(ctx, items)
	if err != nil {
		fmt.Printf("Ошибка обработки: %v\n", err)
		return
	}

	// выводим результаты
	fmt.Println("=== Результаты обработки ===")
	for i, result := range results {
		fmt.Printf("%d: %s\n", i, result)
	}
}
