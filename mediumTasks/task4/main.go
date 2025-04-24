package main

//4я задача

import "fmt"

func main() {
	readers := make(map[string]map[string][]string)

	readers["Иванов"] = map[string][]string{
		"книги":     {"Война и мир", "Преступление и наказание"},
		"периодика": {"Газета 'Правда'", "Журнал 'Наука'"},
	}

	readers["Петрова"] = map[string][]string{
		"книги":     {"1984", "Мастер и Маргарита"},
		"периодика": {"Журнал 'Мода'"},
	}

	fmt.Printf("Количество читателей: %d\n\n", len(readers))

	for reader, publications := range readers {
		total := 0
		for pubType, items := range publications {
			total += len(items)
			fmt.Printf("%s: %s - %d шт.\n", reader, pubType, len(items))
		}
		fmt.Printf("Всего у %s: %d изданий\n\n", reader, total)
	}
}
