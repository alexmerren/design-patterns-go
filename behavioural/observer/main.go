package main

import (
	"fmt"
)

const mailNotificationLimitPence = 10_000_00

func main() {
	ledger := &GlobalLedger{}
	pushNotifier := &PushNotifier{}
	mailNotifier := &MailNotifier{}

	account := NewBankAccount("Frankie")
	account.RegisterObserver(ledger)
	account.RegisterObserver(mailNotifier)
	account.RegisterObserver(pushNotifier)

	account2 := NewBankAccount("Alex")
	account2.RegisterObserver(ledger)
	account2.RegisterObserver(mailNotifier)

	account.Transaction(account2, 10_00)
	account2.Transaction(account, 10_000_00)

	fmt.Printf("%s: £%d\n", account.Name, account.BalancePence)
	fmt.Printf("%s: £%d\n", account2.Name, account2.BalancePence)
}

type Transaction struct {
	AccountTo   *BankAccount
	AccountFrom *BankAccount
	AmountPence int
}

func NewTransaction(accountFrom, accountTo *BankAccount, amountPence int) Transaction {
	accountFrom.BalancePence -= amountPence
	accountTo.BalancePence += amountPence
	return Transaction{
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		AmountPence: amountPence,
	}
}
