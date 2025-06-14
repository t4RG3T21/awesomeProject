package main

import (
	"fmt"
	"sync"
	"time"
)

func safeWorker(id int, wg *sync.WaitGroup, panicChan chan<- string) {
	defer func() {
		if r := recover(); r != nil {
			// Перехватываем панику и отправляем информацию в канал
			panicChan <- fmt.Sprintf("🔥 Горутина %d: ПАНИКА ПЕРЕХВАЧЕНА: %v", id, r)
		}
		wg.Done() // Гарантируем завершение даже при панике
	}()

	fmt.Printf("🚦 Горутина %d: запущена\n", id)

	// Имитация работы
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)

	// Специально вызываем панику в одной из горутин
	if id == 3 {
		panic(fmt.Sprintf("💥 Горутина %d: КРИТИЧЕСКАЯ ОШИБКА!", id))
	}

	fmt.Printf("✅ Горутина %d: успешно завершена\n", id)
}

func main() {
	const numWorkers = 5
	var wg sync.WaitGroup
	panicChan := make(chan string, numWorkers) // Буферизованный канал для паник

	fmt.Println("🏁 Главный поток: запуск горутин")
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go safeWorker(i, &wg, panicChan)
	}

	// Запускаем мониторинг паник в отдельной горутине
	go func() {
		wg.Wait()        // Ждем завершения всех рабочих
		close(panicChan) // Закрываем канал после завершения
	}()

	fmt.Println("👂 Главный поток: ожидание завершения и обработка ошибок...")

	// Обрабатываем все сообщения о панике
	for panicMsg := range panicChan {
		fmt.Println("\n❗️ ОБРАБОТКА ПАНИКИ В ГЛАВНОМ ПОТОКЕ:")
		fmt.Println(panicMsg)
		fmt.Println("🛡️ Приняты меры: восстановление работы системы")
	}

	fmt.Println("\n🟢 Все горутины завершены. Программа работает стабильно.")
}
