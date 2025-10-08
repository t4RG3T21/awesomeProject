package main

import (
	"fmt"
	"sync"
)

//5. Банковский счет с RW Mutex
//Создайте структуру банковского счета, где:
//Проверка баланса происходит очень часто
//Изменение баланса (пополнение/списание) происходит редко
//Оптимизируйте с помощью dirsync.RWMutex.

func main() {
	account := BankAccount{balance: 100.0}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				balance := account.CheckBalance()
				fmt.Printf("Goroutine %d: Balance $%.2f\n", i, balance)
			}
		}()
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		account.Deposit(50.0)
		fmt.Println("Deposit successful(50.0$)")
	}()

	go func() {
		defer wg.Done()
		if success := account.Withdraw(30.0); success {
			fmt.Println("Withdraw successful(30.0$)")
		} else {
			fmt.Println("Withdraw failed(30.0$)")
		}
	}()
	wg.Wait()
	finalBalance := account.CheckBalance()
	fmt.Printf("Final Balance: $%.2f\n", finalBalance)
}

type BankAccount struct {
	mu      sync.RWMutex
	balance float64
}

func (acc *BankAccount) CheckBalance() float64 {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	return acc.balance
}

func (acc *BankAccount) Deposit(amount float64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += amount
}

func (acc *BankAccount) Withdraw(amount float64) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.balance >= amount {
		acc.balance -= amount
		return true
	}

	return false
}
