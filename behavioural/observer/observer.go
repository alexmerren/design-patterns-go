package main

import "fmt"

// ** OBSERVER INTERFACE **

type TransactionObserver interface {
	Notify(Transaction)
}

// ** CONCRETE OBSERVER **

type PushNotifier struct{}

func (p *PushNotifier) Notify(t Transaction) {
	fmt.Printf("Pushed notification to %s for transaction amount %d to %s\n", t.AccountFrom.Name, t.AmountPence, t.AccountTo.Name)
}

type MailNotifier struct{}

func (m *MailNotifier) Notify(t Transaction) {
	if t.AmountPence >= mailNotificationLimitPence {
		fmt.Printf("Mailed %s for transaction amount %d to %s\n", t.AccountFrom.Name, t.AmountPence, t.AccountTo.Name)
	}
}

type GlobalLedger struct {
	Transactions []Transaction
	Accounts     []*BankAccount
}

func (g *GlobalLedger) Notify(t Transaction) {
	g.AddTransaction(t)
}

func (g *GlobalLedger) AddAccount(account *BankAccount) {
	g.Accounts = append(g.Accounts, account)
}

func (g *GlobalLedger) AddTransaction(transaction Transaction) {
	g.Transactions = append(g.Transactions, transaction)
}
