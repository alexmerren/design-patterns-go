package main

import "fmt"

const overdraftLimit = -100_00

type BankAccountReciever interface {
	Deposit(int)
	Withdraw(int) bool
}

type BankAccount struct {
	name         string
	balancePence int
}

func NewBankAccount(name string) *BankAccount {
	return &BankAccount{
		name:         name,
		balancePence: 0,
	}
}

func (b *BankAccount) GetBalance() int {
	return b.balancePence
}

func (b *BankAccount) Deposit(amountPence int) {
	fmt.Printf("Depositing %dp into %s's Account\n", amountPence, b.name)
	b.balancePence += amountPence
}

func (b *BankAccount) Withdraw(amountPence int) bool {
	if b.balancePence-amountPence <= overdraftLimit {
		fmt.Printf("%s's Account is over the overdraft limit\n", b.name)
		return false
	}
	fmt.Printf("Withdrawing %dp from %s's Account\n", amountPence, b.name)
	b.balancePence -= amountPence
	return true
}
