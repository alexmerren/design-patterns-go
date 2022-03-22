package main

import "fmt"

const (
	overdraftLimit   = -100_00
	suspiciousAmount = 10_000_00
)

type ProxyBankAccount struct {
	account Subject
}

func NewProxyBankAccount(name string, balancePence int) *ProxyBankAccount {
	return &ProxyBankAccount{
		account: NewBankAccount(name, balancePence),
	}
}

func (b *ProxyBankAccount) Deposit(amountPence int) int {
	if amountPence >= suspiciousAmount {
		fmt.Println("Woah. That's a lot of money")
	}
	return b.account.Deposit(amountPence)
}

func (b *ProxyBankAccount) Withdraw(amountPence int) int {
	balance := b.account.Withdraw(amountPence)
	if balance <= overdraftLimit {
		// In reality, this would probably be done with command pattern so
		// undo() can be called, but since we know that the withdraw() and
		// deposit() functions are just adding and subtracting from balance,
		// this is okay.
		b.account.Deposit(amountPence)
		return 0
	}
	return balance
}

func (b *ProxyBankAccount) GetBalance() int {
	return b.account.GetBalance()
}
