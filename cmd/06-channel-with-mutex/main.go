package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Bank struct {
	balance int
	mu      *sync.Mutex
}

func (b *Bank) deposit(amount int) {
	b.mu.Lock()
	b.balance += amount
	b.mu.Unlock()
}

func (b *Bank) withdraw(amount int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.balance >= amount {
		b.balance -= amount
		return true
	}
	return false
}

func main() {
	m := &sync.Mutex{}

	bank := &Bank{
		balance: 0,
		mu:      m,
	}

	for i := 0; i < 10000; i++ {
		go bank.deposit(100)
	}

	time.Sleep(2000)

	fmt.Println("Total Amount (before waiting): " + strconv.Itoa(bank.balance))
}
