package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type ConfigManager struct {
	mu     sync.RWMutex
	config map[string]string
}

func NewConfigManager(initial map[string]string) *ConfigManager {
	cfg := make(map[string]string)
	for k, v := range initial {
		cfg[k] = v
	}
	return &ConfigManager{config: cfg}
}

func (cm *ConfigManager) Get(key string) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.config[key], cm.config[key] != ""
}

func (cm *ConfigManager) Reload(newConfig map[string]string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// Создаем глубокую копию
	updated := make(map[string]string, len(newConfig))
	for k, v := range newConfig {
		updated[k] = v
	}

	cm.config = updated
}

func main() {
	// Исходная конфигурация
	initialConfig := map[string]string{
		"theme":   "dark",
		"timeout": "30s",
		"debug":   "false",
	}

	manager := NewConfigManager(initialConfig)

	// Счетчики для статистики
	var readCount atomic.Int32
	var reloadCount atomic.Int32

	// Запускаем 100 читателей
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				// Читаем случайный параметр
				keys := []string{"theme", "timeout", "debug"}
				key := keys[rand.Intn(len(keys))]

				val, ok := manager.Get(key)
				if ok {
					fmt.Printf("Reader %d: %s = %s\n", id, key, val)
				}
				readCount.Add(1)
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			}
		}(i)
	}

	// Запускаем перезагрузчик
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)

			// Генерируем новую конфигурацию
			themes := []string{"dark", "light", "blue"}
			newConfig := map[string]string{
				"theme":   themes[rand.Intn(len(themes))],
				"timeout": fmt.Sprintf("%ds", rand.Intn(60)+1),
				"debug":   fmt.Sprintf("%t", rand.Intn(2) == 0),
			}

			manager.Reload(newConfig)
			reloadCount.Add(1)
			fmt.Printf("\n🔥 Конфигурация перезагружена #%d\n", reloadCount.Load())
		}
	}()

	wg.Wait()
	fmt.Printf("\n✅ Все операции завершены\nЧтений: %d\nПерезагрузок: %d\n",
		readCount.Load(), reloadCount.Load())
}
