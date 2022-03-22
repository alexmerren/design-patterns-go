package main

import "fmt"

type Notification struct {
}

func (n *Notification) debitNotification() {
	fmt.Println("Sending account debit notification")
}
