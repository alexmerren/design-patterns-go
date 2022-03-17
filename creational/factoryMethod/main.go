package main

import "fmt"

func main() {
	bank := NewBank()
	account1 := bank.Create(Personal, 100_00)
	account2 := bank.Create(Saver, 86_000_00)
	fmt.Printf("Balance: %d, Type: %s\n", account1.GetBalance(), account1.GetType())
	fmt.Printf("Balance: %d, Type: %s\n", account2.GetBalance(), account2.GetType())
}

type Bank struct {
	factory *AccountFactory
	name    string
}

func NewBank() *Bank {
	return &Bank{
		factory: NewAccountFactory(),
		name:    "Alex's Bank",
	}
}

func (b *Bank) Create(accountType AccountType, amountPence int) Product {
	account, err := b.factory.Create(accountType, amountPence)
	if err != nil {
		return nil
	}
	return account
}
