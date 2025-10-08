package main

//7. статистических данных, где:
//Данные обновляются раз в минуту (одна горутина)
//Данные читаются сотни раз в секунду (много горутин)
//Примените dirsync.RWMutex.

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	collector := NewStats()

	go func() {
		for {
			NewStats := map[string]float64{
				"cpu_usage":           75.3,
				"mem_usage":           1000.0,
				"requests_per_second": 123.7,
			}
			collector.UpdateStats(NewStats)
			time.Sleep(1 * time.Minute)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				if cpu, exists := collector.GetStats("cpu_usage"); exists {
					fmt.Printf("Goroutine %d: CPU usage is %.1f%%\n", id, cpu)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
}

type Stats struct {
	mu    sync.RWMutex
	stats map[string]float64
}

func NewStats() *Stats {
	return &Stats{
		stats: make(map[string]float64),
	}
}

func (s *Stats) UpdateStats(newStats map[string]float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stats = newStats
	fmt.Println("Stats updated")
}

func (s *Stats) GetStats(key string) (float64, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists := s.stats[key]
	return value, exists
}
