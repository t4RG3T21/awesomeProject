package main

import (
	"fmt"
	"time"
)

//2. Синхронизация горутин
//Напишите программу, где:
//Главная goroutine ждёт завершения двух рабочих goroutine с помощью канала.
//Рабочие goroutine отправляют сигнал (done <- true) после выполнения задачи.
//Цель: научиться синхронизировать горутины.

func worker(id int, done chan<- bool) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println("worker", id, "finished")
	done <- true
}

func main() {
	done := make(chan bool)

	go worker(1, done)
	go worker(2, done)

	<-done
	<-done

	fmt.Println("all workers finished")
}
