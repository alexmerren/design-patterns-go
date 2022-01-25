package main

type Component interface {
	GetBalance() int
}

type BankAccount struct {
	id           string
	balancePence int
}

func NewBankAccount(id string, balancePence int) *BankAccount {
	return &BankAccount{
		id:           id,
		balancePence: balancePence,
	}
}

func (b *BankAccount) GetBalance() int {
	return b.balancePence
}
