package main

import "errors"

type AccountType int

const (
	Personal AccountType = 1
	Saver    AccountType = 2
)

type Creator interface {
	Create(AccountType, int) (*Product, error)
}

type AccountFactory struct{}

func NewAccountFactory() *AccountFactory {
	return &AccountFactory{}
}

func (a *AccountFactory) Create(accountType AccountType, amountPence int) (Product, error) {
	switch accountType {
	case Personal:
		return &PersonalAccount{
			balancePence: amountPence,
			accountType:  "Personal",
		}, nil
	case Saver:
		return &SaverAccount{
			balancePence: amountPence,
			accountType:  "Saver",
		}, nil
	default:
		return nil, errors.New("Type Not Supported")
	}
}
