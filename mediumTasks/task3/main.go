package main

//3я задача

import "fmt"

func main() {
	days := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}

	// копирую рабочие дни
	workDays := make([]string, 5)
	copy(workDays, days[:5]) // пн-пт

	// оставляю в исходном слайсе только выходные
	days = days[5:] // days = (суббота, воскресенье)

	fmt.Println("Рабочие дни:", workDays)
	fmt.Println("Выходные дни:", days)
}
