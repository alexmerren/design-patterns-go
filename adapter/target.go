package main

import "fmt"

type Target interface {
	Send(string, string, int) error
}

type Paypal struct{}

func NewPaypal() *Paypal {
	return &Paypal{}
}

func (p *Paypal) Send(accountFrom, accountTo string, amountPence int) error {
	fmt.Printf("Transfered %d from %s to %s via paypal\n", amountPence, accountFrom, accountTo)
	return nil
}
