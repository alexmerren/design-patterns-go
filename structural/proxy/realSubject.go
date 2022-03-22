package main

type BankAccount struct {
	name         string
	balancePence int
}

func NewBankAccount(name string, initialBalancePence int) *BankAccount {
	return &BankAccount{
		name:         name,
		balancePence: initialBalancePence,
	}
}

func (b *BankAccount) Deposit(amountPence int) int {
	b.balancePence += amountPence
	return b.balancePence
}

func (b *BankAccount) Withdraw(amountPence int) int {
	b.balancePence -= amountPence
	return b.balancePence
}

func (b *BankAccount) GetBalance() int {
	return b.balancePence
}
