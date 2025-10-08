package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Отложенная функция для обработки паники
	defer func() {
		if r := recover(); r != nil {
			// Восстанавливаемся после паники
			fmt.Println("Произошла ошибка:", r)

			// Читаем и выводим содержимое выходного файла
			content, err := os.ReadFile("data/data_out.txt")
			if err != nil {
				fmt.Println("Ошибка при чтении выходного файла:", err)
			} else {
				fmt.Println("\nСодержимое файла data_out.txt до ошибки:")
				fmt.Println(string(content))
			}
		}
	}()

	// Создаем директорию data, если её нет
	if err := os.MkdirAll("data", 0755); err != nil {
		panic("не удалось создать директорию data: " + err.Error())
	}

	// Открываем входной файл
	inFile, err := os.Open("in1.txt")
	if err != nil {
		panic("не удалось открыть файл in1.txt: " + err.Error())
	}
	defer inFile.Close()

	// Создаем выходной файл
	outFile, err := os.Create("data/data_out.txt")
	if err != nil {
		panic("не удалось создать файл data_out.txt: " + err.Error())
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Читаем входной файл построчно
	scanner := bufio.NewScanner(inFile)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Разбиваем строку на поля
		parts := strings.Split(line, "|")

		// Проверяем что есть 3 поля и все заполнены
		if len(parts) < 3 {
			panic(fmt.Sprintf("parse error: empty field on string %d", lineNumber))
		}

		name := strings.TrimSpace(parts[0])
		address := strings.TrimSpace(parts[1])
		city := strings.TrimSpace(parts[2])

		if name == "" || address == "" || city == "" {
			panic(fmt.Sprintf("parse error: empty field on string %d", lineNumber))
		}

		// Записываем данные в выходной файл
		_, err := writer.WriteString(
			fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n",
				lineNumber, name, address, city))

		if err != nil {
			panic("ошибка записи в файл: " + err.Error())
		}
	}

	// Проверяем ошибки сканера
	if err := scanner.Err(); err != nil {
		panic("ошибка чтения файла: " + err.Error())
	}

	fmt.Println("Обработка файла успешно завершена!")
}
