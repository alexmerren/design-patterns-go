package main

// ** SUBJECT INTERFACE**

type TransactionSubject interface {
	RegisterObserver(TransactionObserver)
	UnregisterObserver(TransactionObserver)
	NotifyObservers()
}

// ** CONCRETE SUBJECT **

type BankAccount struct {
	Name         string
	BalancePence int
	Observers    []TransactionObserver
}

func NewBankAccount(name string) *BankAccount {
	return &BankAccount{
		Name:         name,
		BalancePence: 0,
	}
}

func (a *BankAccount) RegisterObserver(o TransactionObserver) {
	a.Observers = append(a.Observers, o)
}

func (a *BankAccount) UnregisterObserver(o TransactionObserver) {
	a.Observers = removeObserver(a.Observers, o)
}

func (a *BankAccount) NotifyObservers(t Transaction) {
	for _, observer := range a.Observers {
		observer.Notify(t)
	}
}

func (a *BankAccount) Transaction(to *BankAccount, amountPence int) {
	t := NewTransaction(a, to, amountPence)
	a.NotifyObservers(t)
}

func removeObserver(observersList []TransactionObserver, o TransactionObserver) []TransactionObserver {
	for index, observer := range observersList {
		if o == observer {
			tmp := observersList[len(observersList)-1]
			observersList[len(observersList)-1] = observer
			observersList[index] = tmp
			return observersList[:len(observersList)-1]
		}
	}
	return observersList
}
