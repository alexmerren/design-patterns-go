package main

type Caller struct {
	commandList []*BankAccountCommand
}

func NewCaller() *Caller {
	return &Caller{
		commandList: make([]*BankAccountCommand, 0),
	}
}

func (c *Caller) ExecuteCommand(reciever BankAccountReciever, commandType CommandType, amountPence int) {
	command := &BankAccountCommand{
		reciever:    reciever,
		commandType: commandType,
		amountPence: amountPence,
	}
	command.Do()
	if !command.succeded {
		command.Undo()
	}
}

func (c *Caller) StoreCommand(reciever BankAccountReciever, commandType CommandType, amountPence int) {
	command := &BankAccountCommand{
		reciever:    reciever,
		commandType: commandType,
		amountPence: amountPence,
	}
	c.commandList = append(c.commandList, command)
}

func (c *Caller) ExecuteCommands() {
	for _, command := range c.commandList {
		command.Do()
		if !command.succeded {
			command.Undo()
		}
	}
}
