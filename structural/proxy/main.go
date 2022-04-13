package main

func main() {
	// This is the real subject - no access control, just functions to add and
	// subtract from a balance in a bank account.
	account1 := NewBankAccount("Alex's insecure bank account", 100_00)
	account1.Deposit(100_00)

	// This is the proxy - a bank account that we have added access control
	// onto. This will now check if a withdrawal can be possible and doesn't
	// leave them below the overdraft limit. The proxy can also notify relevant
	// people if a withdrawal/deposit is above a suspicious amount (to be
	// determined by the business).
	account2 := NewProxyBankAccount("Alex's secure bank account", 1_234_00)
	account2.Withdraw(1_000_000_00)

	// This is the result of someone hitting that suspicious amount - but for
	// Bill Gates, this is nothing.
	account3 := NewProxyBankAccount("Bill Gates", 90_000_000_000_00)
	account3.Withdraw(12_728_193_00)
}
