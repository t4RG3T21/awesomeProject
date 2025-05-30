package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Привет, я %s, мне %d лет\n", p.Name, p.Age)
}

type Employee struct {
	Person Person // Встраивание (не указываем имя поля)
	Salary float64
}

func (e Employee) ShowSalary() {
	fmt.Printf("Моя зарплата: $%.2f\n", e.Salary)
}

func main() {
	// Создаем сотрудника
	emp := Employee{
		Person: Person{
			Name: "Анна",
			Age:  28,
		},
		Salary: 75000.50,
	}

	fmt.Println("Имя:", emp.Name)    // Анна
	fmt.Println("Возраст:", emp.Age) // 28
	emp.Introduce()                  // Привет, я Анна, мне 28 лет

	fmt.Println("Зарплата:", emp.Salary) // 75000.50
	emp.ShowSalary()                     // Моя зарплата: $75000.50

}
