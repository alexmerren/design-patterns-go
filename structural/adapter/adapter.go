package main

import "math/rand"

type BankAdapter struct {
	Bank *Bank
}

func NewBankAdapter(bank *Bank) *BankAdapter {
	return &BankAdapter{
		Bank: bank,
	}
}

func (b *BankAdapter) Send(emailFrom, emailTo string, amountPence int) error {
	accountFrom := b.findAccountByEmail(emailFrom)
	accountTo := b.findAccountByEmail(emailTo)
	return b.Bank.ProcessPayment(accountFrom, accountTo, amountPence)
}

func (b *BankAdapter) findAccountByEmail(email string) int {
	return rand.Intn(1000000)
}
