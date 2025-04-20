package main

//2 часть 1й задачи
import (
	"fmt"
	"math"
)

type EuropeanVelocity float64
type AmericanVelocity float64

func main() {
	// 120.4 м/с в км/ч (1 м/с = 3.6 км/ч)
	mpsToKph := 120.4 * 3.6
	ev := EuropeanVelocity(math.Round(mpsToKph*100) / 100) // округление до 2 знаков

	// 130 м/с в миль/ч (1 м/с = +- 2.23694 миль/ч)
	mpsToMph := 130 * 2.23694
	av := AmericanVelocity(math.Round(mpsToMph*100) / 100) // округление

	fmt.Println("Европейская скорость:", ev, "км/ч")
	fmt.Println("Американская скорость:", av, "миль/ч")
}
