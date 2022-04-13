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

func (b *BankAccount) Deposit(amountPence int) (int, error) {
	b.balancePence += amountPence
	return b.balancePence, nil
}

func (b *BankAccount) Withdraw(amountPence int) (int, error) {
	b.balancePence -= amountPence
	return b.balancePence, nil
}

func (b *BankAccount) GetBalance() (int, error) {
	return b.balancePence, nil
}

func (b *BankAccount) GetName() string {
	return b.name
}
