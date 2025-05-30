package main

import (
	"fmt"
	"time"
)

type Event struct {
	Title string
	Time  time.Time
}

func (e Event) IsAfter(now time.Time) bool {
	return e.Time.After(now)
}

func main() {
	futureEvent := Event{
		Title: "Защита диплома",
		Time:  time.Date(2024, time.June, 15, 10, 0, 0, 0, time.UTC),
	}

	pastEvent := Event{
		Title: "Знакомство с Go",
		Time:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	currentTime := time.Now()

	fmt.Printf("Событие '%s' еще не наступило? %t\n",
		futureEvent.Title, futureEvent.IsAfter(currentTime))

	fmt.Printf("Событие '%s' еще не наступило? %t\n",
		pastEvent.Title, pastEvent.IsAfter(currentTime))
}
