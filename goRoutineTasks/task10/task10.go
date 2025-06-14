package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func worker(stopFlag *atomic.Bool, id int) {
	fmt.Printf("🏁 Рабочий %d: начал работу\n", id)

	// Бесконечный цикл с проверкой флага
	for {
		// Периодически проверяем флаг завершения
		if stopFlag.Load() {
			fmt.Printf("🛑 Рабочий %d: получен сигнал остановки\n", id)
			return
		}

		// Выполняем полезную работу
		fmt.Printf("🔧 Рабочий %d: выполняет задачу...\n", id)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var stopFlag atomic.Bool // Атомарный флаг остановки

	// Запускаем 3 рабочих
	for i := 1; i <= 3; i++ {
		go worker(&stopFlag, i)
	}

	// Даем горутинам поработать 3 секунды
	fmt.Println("👨‍💼 Менеджер: рабочие начали, жду 3 секунды...")
	time.Sleep(3 * time.Second)

	// Устанавливаем флаг остановки
	fmt.Println("\n👨‍💼 Менеджер: отправляю сигнал остановки всем рабочим!")
	stopFlag.Store(true)

	// Даем время для корректного завершения
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\n✅ Менеджер: все рабочие завершили работу")
}
