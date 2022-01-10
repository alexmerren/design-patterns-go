package main

import "fmt"

type Adaptee interface {
	ProcessPayment(int, int, int) error
}

type Bank struct{}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) ProcessPayment(accountFrom, accountTo int, amountPence int) error {
	fmt.Printf("Transfered %d from account %d to account %d via bank transfer\n", amountPence, accountFrom, accountTo)
	return nil
}
