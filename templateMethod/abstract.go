package main

import "errors"

const OVERDRAFT_LIMIT = -100_00

type Template interface {
	ProcessTransaction(int) error
	GetAccountBalance() int
	ChangeAccountBalance(int) error
	SendNotification()
	AppendToLedger(int) error
}

type Transaction struct {
	t Template
}

func NewTransaction(t Template) *Transaction {
	return &Transaction{
		t: t,
	}
}

func (t *Transaction) ProcessTransaction(amountPence int) error {
	if t.t.GetAccountBalance()-amountPence < OVERDRAFT_LIMIT {
		return errors.New("Not enough balance")
	}

	if err := t.t.ChangeAccountBalance(amountPence); err != nil {
		return errors.New("Could not change balance of account")
	}

	t.t.SendNotification()

	if err := t.t.AppendToLedger(amountPence); err != nil {
		return errors.New("Failed to send transaction to Ledger")
	}

	return nil
}
