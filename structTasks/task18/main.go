package main

import "fmt"

type Transport interface {
	Move() string
}

type Car struct {
	Model string
}

func (c Car) Move() string {
	return fmt.Sprintf("%s едет по дороге", c.Model)
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() string {
	return fmt.Sprintf("%s велосипед крутит педали", b.Brand)
}

type Plane struct {
	FlightNumber string
}

func (p Plane) Move() string {
	return fmt.Sprintf("Самолет %s взлетает в небо", p.FlightNumber)
}

func StartRace(transports []Transport) {
	for i, t := range transports {
		fmt.Printf("Участник %d: %s\n", i+1, t.Move())
	}
}

func main() {
	vehicles := []Transport{
		Car{"Tesla Model S"},
		Bicycle{"Stels"},
		Plane{"SU-123"},
		Car{"Lada Vesta"},
		Bicycle{"Forward"},
	}

	fmt.Println("=== НАЧАЛО ГОНКИ ===")
	StartRace(vehicles)
	fmt.Println("=== ГОНКА ЗАВЕРШЕНА ===")
}
