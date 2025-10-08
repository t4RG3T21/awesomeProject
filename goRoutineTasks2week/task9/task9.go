package main

import (
	"fmt"
	"sync"
	"time"
)

type TimestampStore struct {
	mu         sync.RWMutex
	lastEvents map[string]time.Time
}

func NewTimestampStore() *TimestampStore {
	return &TimestampStore{
		lastEvents: make(map[string]time.Time),
	}
}

func (ts *TimestampStore) UpdateTimestamp(eventName string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.lastEvents[eventName] = time.Now()
}

func (ts *TimestampStore) GetLastTimestamp(eventName string) (time.Time, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	timestamp, exists := ts.lastEvents[eventName]
	return timestamp, exists
}

func main() {
	store := NewTimestampStore()

	store.UpdateTimestamp("user_login")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				if ts, exists := store.GetLastTimestamp("user_login"); exists {
					fmt.Printf("Reader %d-%d: Last login was at %v\n", id, j, ts)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
}
