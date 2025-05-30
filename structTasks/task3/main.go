package main

import "fmt"

func main() {
	counter := Counter{value: 5}

	fmt.Println(counter.GetValue())

	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println(counter.GetValue())

}

type Counter struct {
	value int
}

func (c *Counter) GetValue() float64 {
	return float64(c.value)
}

func (c *Counter) Increment() {
	c.value++
}
