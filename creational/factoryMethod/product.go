package main

type Product interface {
	GetBalance() int
	GetType() string
}

type PersonalAccount struct {
	balancePence int
	accountType  string
}

func (a *PersonalAccount) GetBalance() int {
	return a.balancePence
}

func (a *PersonalAccount) GetType() string {
	return a.accountType
}

type SaverAccount struct {
	balancePence int
	accountType  string
}

func (a *SaverAccount) GetBalance() int {
	return a.balancePence
}

func (a *SaverAccount) GetType() string {
	return a.accountType
}
