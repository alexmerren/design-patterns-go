package main

type Subject interface {
	Deposit(int) int
	Withdraw(int) int
	GetBalance() int
}
