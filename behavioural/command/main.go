package main

import "fmt"

func main() {
	account1 := NewBankAccount("Alex")
	caller := NewCaller()

	caller.ExecuteCommand(account1, CommandTypeDeposit, 1_190_00)
	caller.StoreCommand(account1, CommandTypeDeposit, 100_00)
	caller.StoreCommand(account1, CommandTypeDeposit, 100_00)
	caller.StoreCommand(account1, CommandTypeWithdraw, 100_00)
	caller.StoreCommand(account1, CommandTypeWithdraw, 100_00)
	caller.StoreCommand(account1, CommandTypeWithdraw, 100_00)
	caller.ExecuteCommands()

	fmt.Println(account1.GetBalance())
}
