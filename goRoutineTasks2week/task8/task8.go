package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ShutdownManager struct {
	flag atomic.Bool
}

func (sm *ShutdownManager) IsShutdown() bool {
	return sm.flag.Load()
}

func (sm *ShutdownManager) Shutdown() {
	sm.flag.Store(true)
}

func main() {
	var manager ShutdownManager
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for !manager.IsShutdown() {
				fmt.Printf("Goroutine %d: working...\n", id)
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Printf("Goroutine %d: shutting down gracefully\n", id)
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Main: initiating shutdown...")
	manager.Shutdown()

	wg.Wait()
	fmt.Println("Main: done")
}
