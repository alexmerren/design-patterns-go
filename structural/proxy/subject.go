package main

type Subject interface {
	Deposit(int) (int, error)
	Withdraw(int) (int, error)
	GetBalance() (int, error)
	GetName() string
}
