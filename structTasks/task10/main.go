package main

import (
	"errors"
	"fmt"
)

// Ошибка для случая недостаточного баланса

var ErrInsufficientFunds = errors.New("недостаточно средств на счете")

// Структура кошелька

type Wallet struct {
	Balance float64
}

// Метод для списания средств

func (w *Wallet) Spend(amount float64) error {
	if amount > w.Balance {
		return ErrInsufficientFunds
	}
	w.Balance -= amount
	return nil
}

func main() {
	// Создаем кошелек с балансом 100.0
	wallet := &Wallet{Balance: 100.0}

	// Пытаемся потратить 80
	if err := wallet.Spend(80); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Списание 80. Новый баланс: %.2f\n", wallet.Balance)
	}

	// Пытаемся потратить 50
	if err := wallet.Spend(50); err != nil {
		fmt.Println("Ошибка:", err) // Сработает эта ветка
	} else {
		fmt.Printf("Списание 50. Новый баланс: %.2f\n", wallet.Balance)
	}
}
