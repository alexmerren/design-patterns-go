package main

type Facade interface {
	WithdrawCash(string, int, int) error
}

type ATMFacade struct {
	account       *Account
	securityCode  *SecurityCode
	notification  *Notification
	ledger        *Ledger
	counterLedger *Ledger
}

func NewATMFacade(accountID string, pinNumber int) *ATMFacade {
	return &ATMFacade{
		account:       NewAccount(accountID),
		securityCode:  NewSecurityCode(pinNumber),
		notification:  &Notification{},
		ledger:        &Ledger{},
		counterLedger: &Ledger{},
	}
}

func (f *ATMFacade) WithdrawCash(accountID string, pinNumber int, amountPence int) error {
	err := f.account.checkAccountID(accountID)
	if err != nil {
		return err
	}

	err = f.securityCode.checkCode(pinNumber)
	if err != nil {
		return err
	}

	err = f.account.withdraw(amountPence)
	if err != nil {
		return err
	}

	f.notification.debitNotification()
	f.ledger.makeEntry(accountID, "debit", amountPence)
	f.counterLedger.makeEntry(accountID, "credit", amountPence)
	return nil
}
