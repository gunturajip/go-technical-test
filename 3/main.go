package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	ID      string
	Balance float64
	mu      sync.Mutex
}

func NewAccount(id string, initialBalance float64) *Account {
	return &Account{
		ID:      id,
		Balance: initialBalance,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance < amount {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.Balance
}

type Bank struct {
	accounts map[string]*Account
	mu       sync.RWMutex
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[string]*Account),
	}
}

func (b *Bank) CreateAccount(accountID string, initialBalance float64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.accounts[accountID]; exists {
		return fmt.Errorf("account %s already exists", accountID)
	}

	b.accounts[accountID] = NewAccount(accountID, initialBalance)
	return nil
}

func (b *Bank) GetAccount(accountID string) (*Account, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	account, exists := b.accounts[accountID]
	if !exists {
		return nil, fmt.Errorf("account %s not found", accountID)
	}

	return account, nil
}

func (b *Bank) Transfer(fromID, toID string, amount float64) error {
	if fromID == toID {
		return errors.New("cannot transfer to the same account")
	}

	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	b.mu.Lock()
	fromAccount, exists := b.accounts[fromID]
	if !exists {
		b.mu.Unlock()
		return fmt.Errorf("source account %s not found", fromID)
	}

	toAccount, exists := b.accounts[toID]
	if !exists {
		b.mu.Unlock()
		return fmt.Errorf("destination account %s not found", toID)
	}
	b.mu.Unlock()

	if fromID < toID {
		fromAccount.mu.Lock()
		toAccount.mu.Lock()
	} else {
		toAccount.mu.Lock()
		fromAccount.mu.Lock()
	}
	defer fromAccount.mu.Unlock()
	defer toAccount.mu.Unlock()

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds in source account")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount
	return nil
}

func main() {
	bank := NewBank()

	bank.CreateAccount("acc1", 1000.0)
	bank.CreateAccount("acc2", 500.0)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc, _ := bank.GetAccount("acc1")
			acc.Deposit(100.0)
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc, _ := bank.GetAccount("acc2")
			acc.Withdraw(50.0)
		}()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bank.Transfer("acc1", "acc2", 75.0)
		}()
	}

	wg.Wait()

	acc1, _ := bank.GetAccount("acc1")
	acc2, _ := bank.GetAccount("acc2")
	fmt.Printf("Account %s balance: %.2f\n", acc1.ID, acc1.GetBalance())
	fmt.Printf("Account %s balance: %.2f\n", acc2.ID, acc2.GetBalance())
}
