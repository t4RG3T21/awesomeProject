package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
	}

	for _, shape := range shapes {
		fmt.Printf(
			"Тип: %T\nПлощадь: %.2f\nПериметр: %.2f\n\n",
			shape,
			shape.Area(),
			shape.Perimeter(),
		)
	}
}
