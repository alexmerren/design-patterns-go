package main

import "fmt"

type Ledger struct{}

func (l *Ledger) makeEntry(accountID, transactionType string, amountPence int) {
	fmt.Printf(
		"Ledger entry for accountID %s with type %s for amount %d\n",
		accountID, transactionType, amountPence)
}
