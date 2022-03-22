package main

import (
	"errors"
	"fmt"
)

const BALANCE_LIMIT = -100_00

type Account struct {
	id           string
	balancePence int
}

func NewAccount(id string) *Account {
	return &Account{
		id:           id,
		balancePence: 0,
	}
}

func (a *Account) checkAccountID(id string) error {
	if a.id != id {
		return errors.New("Account ID is incorrect")
	}
	return nil
}

func (a *Account) withdraw(amountPence int) error {
	if a.balancePence-amountPence < BALANCE_LIMIT {
		return errors.New("Insufficient account funds")
	}
	a.balancePence -= amountPence
	fmt.Printf("Withdrew %d from account id %s\n", amountPence, a.id)
	return nil
}
