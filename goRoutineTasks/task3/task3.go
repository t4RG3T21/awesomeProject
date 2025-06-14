package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Помощник: Отправляюсь в магазин...")
		time.Sleep(2 * time.Second)
		fmt.Println("Помощник: Done! Вернулся с покупками")
	}()

	fmt.Println("Главный: Жду помощника...")
	wg.Wait()
	fmt.Println("Главный: Ура! Можем начинать готовку")
}
