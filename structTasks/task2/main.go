package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}

	fmt.Println("Площадь прямоугольника:", rect.Area())
	fmt.Println("Площадь круга:", circle.Area())

}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
