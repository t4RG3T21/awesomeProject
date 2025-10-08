package main

import (
	"fmt"
	"sync"
	"time"
)

// worker - функция, которая будет выполняться в каждой горутине-воркере
// Параметры:
//
//	id - идентификатор воркера (для наглядности)
//	jobs - канал для получения задач (только для чтения)
//	results - канал для отправки результатов (только для записи)
//	wg - WaitGroup для отслеживания завершения работы
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы

	// Цикл по задачам из канала jobs
	// Цикл завершится, когда канал jobs будет закрыт и все задачи обработаны
	for job := range jobs {
		fmt.Printf("Воркер %d начал обработку задачи %d\n", id, job)

		// Имитируем обработку задачи (например, сложные вычисления)
		time.Sleep(1 * time.Second)

		// Выполняем работу: умножаем число на 2
		result := job * 2

		//Отправляем результат в канал results
		results <- result
		fmt.Printf("Воркер %d завершил задачу %d, результат: %d\n", id, job, result)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func main() {
	// Количество воркеров в пуле
	const numWorkers = 3
	// Количество задач для обработки
	const numJobs = 10

	// Создаем каналы для задач и результатов
	jobs := make(chan int, numJobs)    // Буферизованный канал для задач
	results := make(chan int, numJobs) // Буферизованный канал для результатов

	// Создаем WaitGroup для отслеживания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup для каждого воркера
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задачи в канал jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j // Отправляем номер задачи
	}
	close(jobs) // Закрываем канал задач после отправки всех зданий

	// Ждем завершения всех воркеров
	wg.Wait()
	close(results) // Закрываем канал результатов после завершения

	// Собираем и выводим результаты
	fmt.Println("\n === Результаты ===")
	for result := range results {
		fmt.Printf("Результат: %d\n", result)
	}
}
