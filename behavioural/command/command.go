package main

type CommandType int

const (
	CommandTypeDeposit  CommandType = 1
	CommandTypeWithdraw CommandType = 2
)

type BankAccountCommander interface {
	Do()
	Undo()
}

type BankAccountCommand struct {
	reciever    BankAccountReciever
	commandType CommandType
	amountPence int
	succeded    bool
}

func (b *BankAccountCommand) Do() {
	switch b.commandType {
	case CommandTypeDeposit:
		b.reciever.Deposit(b.amountPence)
		b.succeded = true
	case CommandTypeWithdraw:
		b.succeded = b.reciever.Withdraw(b.amountPence)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeded {
		return
	}
	switch b.commandType {
	case CommandTypeDeposit:
		b.reciever.Withdraw(b.amountPence)
	case CommandTypeWithdraw:
		b.reciever.Deposit(b.amountPence)
	}
}
