package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicBool struct {
	flag uint32
}

func (b *AtomicBool) Set(value bool) {
	var val uint32 = 0
	if value {
		val = 1
	}
	atomic.StoreUint32(&b.flag, val)
}

func (b *AtomicBool) Get() bool {
	return atomic.LoadUint32(&b.flag) == 1
}

func (b *AtomicBool) SetTrue() {
	atomic.StoreUint32(&b.flag, 1)
}

func (b *AtomicBool) SetFalse() {
	atomic.StoreUint32(&b.flag, 0)
}

func (b *AtomicBool) TrySetTrue() bool {
	return atomic.CompareAndSwapUint32(&b.flag, 0, 1)
}

func (b *AtomicBool) TrySetFalse() bool {
	return atomic.CompareAndSwapUint32(&b.flag, 1, 0)
}

func main() {
	var flag AtomicBool
	var wg sync.WaitGroup

	// Запускаем 10 горутин, которые пытаются установить флаг
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if flag.TrySetTrue() {
				fmt.Printf("Горутина %d: установила флаг в true\n", id)
				time.Sleep(100 * time.Millisecond) // Имитация работы
				flag.SetFalse()
			} else {
				fmt.Printf("Горутина %d: флаг уже установлен\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\nФинальное значение флага:", flag.Get())
}
