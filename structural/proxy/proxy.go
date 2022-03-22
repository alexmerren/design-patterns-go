package main

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
	// We should probably tell someone that a suspicious amount of money is being deposited
	if amountPence >= suspiciousAmount {
	}
	return b.account.Deposit(amountPence)
}

func (b *ProxyBankAccount) Withdraw(amountPence int) int {
	// We should probably tell someone that a suspicious amount of money is being withdrawn
	if amountPence >= suspiciousAmount {
	}

	// In reality, this would probably be done with command pattern so
	// undo() can be called, but since we know that the withdraw() and
	// deposit() functions are just adding and subtracting from balance,
	// this is okay.
	balance := b.account.Withdraw(amountPence)
	if balance <= overdraftLimit {
		b.account.Deposit(amountPence)
		return 0
	}
	return balance
}

func (b *ProxyBankAccount) GetBalance() int {
	return b.account.GetBalance()
}
