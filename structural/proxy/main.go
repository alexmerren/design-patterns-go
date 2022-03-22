package main

import "fmt"

func main() {
	account1 := NewBankAccount("Alex", 100_00)
	account1.Deposit(100_00)
	fmt.Println(account1.GetBalance())

	account2 := NewProxyBankAccount("Alex, but protected", 1_234_00)
	account2.Withdraw(1_000_000_00)
	fmt.Println(account2.GetBalance())
}
