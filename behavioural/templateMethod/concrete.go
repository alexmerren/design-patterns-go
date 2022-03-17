package main

import "fmt"

type BankAccount struct {
	Transaction
	id           string
	balancePence int
}

func NewBankAccount(id string, balancePence int) *BankAccount {
	return &BankAccount{
		id:           id,
		balancePence: balancePence,
	}
}

func (b *BankAccount) GetAccountBalance() int {
	// In reality, this might be a database call.
	return b.balancePence
}

func (b *BankAccount) ChangeAccountBalance(amountPence int) error {
	b.balancePence += amountPence
	return nil
}

func (b *BankAccount) SendNotification() {
	fmt.Printf("[BANK] Sending an email to account: %s that balance is now: %d\n", b.id, b.balancePence)
}

func (b *BankAccount) AppendToLedger(amountPence int) error {
	fmt.Printf("[BANK] Sending transaction to ledger and counter-ledger in the bank....\n")
	return nil
}

type PaypalAccount struct {
	Transaction
	email        string
	balancePence int
}

func NewPaypalAccount(email string, balancePence int) *PaypalAccount {
	return &PaypalAccount{
		email:        email,
		balancePence: balancePence,
	}
}

func (b *PaypalAccount) GetAccountBalance() int {
	// In reality, this would be an API call.
	return b.balancePence
}

func (p *PaypalAccount) ChangeAccountBalance(amountPence int) error {
	p.balancePence += amountPence
	return nil
}

func (b *PaypalAccount) SendNotification() {
	fmt.Printf("[PAYPAL] Sending an SMS to %s that balance is now %d\n", b.email, b.balancePence)
}

func (b *PaypalAccount) AppendToLedger(amountPence int) error {
	fmt.Printf("[PAYPAL] Sending transaction to ledger in Paypal...\n")
	return nil
}
