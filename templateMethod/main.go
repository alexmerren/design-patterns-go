package main

import "fmt"

func main() {
	bankAccount := NewBankAccount("1", 100_00)
	bankTransaction := NewTransaction(bankAccount)
	err := bankTransaction.ProcessTransaction(-10_00)
	if err != nil {
		fmt.Printf("%w\n", err)
	}

	paypalAccount := NewPaypalAccount("alex@merren.com", 0_00)
	paypalTransaction := NewTransaction(paypalAccount)
	err = paypalTransaction.ProcessTransaction(10_00)
	if err != nil {
		fmt.Printf("%w\n", err)
	}
}
