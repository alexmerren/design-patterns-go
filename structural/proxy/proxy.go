package main

import (
	"fmt"
)

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

func (b *ProxyBankAccount) Deposit(amountPence int) (int, error) {
	if amountPence >= suspiciousAmount {
		fmt.Println("The amount requested is suspicious...")
	}
	return b.account.Deposit(amountPence)
}

func (b *ProxyBankAccount) Withdraw(amountPence int) (int, error) {
	if amountPence >= suspiciousAmount {
		fmt.Println("The amount requested is suspicious...")
	}

	return b.account.Withdraw(amountPence)
}

func (b *ProxyBankAccount) GetBalance() (int, error) {
	return b.account.GetBalance()
}

func (b *ProxyBankAccount) GetName() string {
	return b.account.GetName()
}
