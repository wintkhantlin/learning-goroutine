package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	balance int
	wg      *sync.WaitGroup
}

func (b *Bank) deposit(amount int) {
	b.balance += amount
	b.wg.Done()
}

func (b *Bank) withdraw(amount int) bool {
	defer b.wg.Done()
	if b.balance >= amount {
		b.balance -= amount
		return true
	}
	return false
}

func main() {
	wg := &sync.WaitGroup{}

	b := &Bank{
		balance: 0,
		wg:      wg,
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		b.deposit(1)
	}

	wg.Wait()

	fmt.Println(b.balance)
}
