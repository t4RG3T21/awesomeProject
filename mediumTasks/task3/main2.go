package main

//3я задача 2я часть

import "fmt"

func main() {
	// слайсы из предыдущей задачи
	weekend := []string{"Суббота", "Воскресенье"}
	work := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница"}

	// объединяем рабочие дни и выходные
	allDays := append(work, weekend...)

	fmt.Println("Все дни недели:", allDays)
}
