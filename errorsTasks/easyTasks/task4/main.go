package main

import (
	"errors"
	"fmt"
	"time"
)

/// РЕШЕНИЕ С ЧАТГПТ, ПОТОМУ ЧТО НЕ ПОНЯЛ КАК РЕШАТЬ, САМО РЕШЕНИЕ РАЗОБРАЛ, НО ВСЁ РАВНО СЛАБО ПОНЯЛ

// Кастомные ошибки для каждой функции
var (
	ErrConnectionFailed = errors.New("ошибка подключения к базе данных")
	ErrUserNotFound     = errors.New("пользователь не найден")
	ErrInvalidOperation = errors.New("недопустимая операция")
)

// Функция 1: Подключение к базе данных
func connectDB() (string, error) {
	fmt.Println("Попытка подключения к БД...")
	// Симулируем случайную ошибку подключения
	if time.Now().Unix()%2 == 0 { // 50% вероятность ошибки
		return "", ErrConnectionFailed
	}
	return "db-connection-id", nil
}

// Функция 2: Получение пользователя
func getUser(dbConn string) (string, error) {
	fmt.Printf("Поиск пользователя через соединение %s...\n", dbConn)
	// Симулируем ошибку поиска пользователя
	if len(dbConn) > 10 { // Всегда true в нормальных условиях
		return "", ErrUserNotFound
	}
	return "user-123", nil
}

// Функция 3: Выполнение операции
func performOperation(userID string) error {
	fmt.Printf("Выполнение операции для пользователя %s...\n", userID)
	// Симулируем ошибку операции
	if userID == "user-123" {
		return ErrInvalidOperation
	}
	return nil
}

// Основная функция, обрабатывающая всю цепочку
func process() error {
	// Шаг 1: Подключение к БД
	dbConn, err := connectDB()
	if err != nil {
		return fmt.Errorf("не удалось начать процесс: %w", err)
	}

	// Шаг 2: Получение пользователя
	userID, err := getUser(dbConn)
	if err != nil {
		return fmt.Errorf("ошибка получения данных: %w", err)
	}

	// Шаг 3: Выполнение операции
	if err := performOperation(userID); err != nil {
		return fmt.Errorf("ошибка выполнения операции: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("=== Запуск системы ===")

	if err := process(); err != nil {
		fmt.Println("\n⚠️ Обнаружена ошибка в процессе выполнения:")

		// Обработка разных типов ошибок
		switch {
		case errors.Is(err, ErrConnectionFailed):
			fmt.Println("  Проблема: Невозможно подключиться к базе данных")
			fmt.Println("  Решение: Проверьте сетевое подключение и параметры БД")

		case errors.Is(err, ErrUserNotFound):
			fmt.Println("  Проблема: Запрошенный пользователь не существует")
			fmt.Println("  Решение: Проверьте идентификатор пользователя")

		case errors.Is(err, ErrInvalidOperation):
			fmt.Println("  Проблема: Операция не может быть выполнена")
			fmt.Println("  Решение: Выберите другую операцию или проверьте права")

		default:
			fmt.Println("  Неизвестная ошибка:", err)
		}

		// Дополнительная информация для разработчика
		fmt.Printf("\nТехническая информация: %v\n", err)
	} else {
		fmt.Println("✅ Процесс успешно завершен!")
	}

	fmt.Println("=== Завершение работы ===")
}
