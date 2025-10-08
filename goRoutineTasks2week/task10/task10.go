package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type AtomicRateLimiter struct {
	counter    atomic.Int64
	limit      int64
	resetTimer *time.Ticker
}

func NewAtomicRateLimiter(limit int64, interval time.Duration) *AtomicRateLimiter {
	rl := &AtomicRateLimiter{
		limit:      limit,
		resetTimer: time.NewTicker(interval),
	}

	go func() {
		for range rl.resetTimer.C {
			rl.counter.Store(0)
			fmt.Println("Counter reset")
		}
	}()
	return rl
}

func (rl *AtomicRateLimiter) Allow() bool {
	current := rl.counter.Add(1)
	return current <= rl.limit
}

func (rl *AtomicRateLimiter) Stop() {
	rl.resetTimer.Stop()
}

func main() {
	limiter := NewAtomicRateLimiter(5, time.Second)
	defer limiter.Stop()

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Allow: ", i)
		} else {
			fmt.Println("Deny: ", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
