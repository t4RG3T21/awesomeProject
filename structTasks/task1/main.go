package main

import "fmt"

func main() {
	rect := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}

	fmt.Printf("Площадь прямоугольника: %.2f\n", rect.Area())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
