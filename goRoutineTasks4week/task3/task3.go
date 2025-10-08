package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//3. Управление жизненным циклом горутин
//Создайте программу, которая:
//Запускает N рабочих горутин (например, 3)
//Каждая горутина периодически выполняет свою задачу
//При получении сигнала SIGINT (Ctrl+C):
//Корректно останавливает все горутины
//Даёт им время на завершение текущих операций
//Выводит статус завершения

func worker(id int, wg *sync.WaitGroup, stopChan <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Printf("Worker %d stopping...\n", id)
			return
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg, stopChan)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	<-sigChan       // Ожидание сигнала
	close(stopChan) // Остановка всех горутин

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("All workers have stopped.")
}
