package main

const OVERDRAFT_LIMIT = -100_00
const OPENING_CHARGE = 300

type InterestCharges struct {
	component Component
}

func (i *InterestCharges) GetBalance() int {
	return i.component.GetBalance() + int(float64(i.component.GetBalance())*0.0001)
}

type OverdraftCharges struct {
	component Component
}

func (o *OverdraftCharges) GetBalance() int {
	balance := o.component.GetBalance()
	if balance < OVERDRAFT_LIMIT {
		return balance + int(float64(balance-OVERDRAFT_LIMIT)*1.399)
	}
	return balance
}

type OpeningCharges struct {
	component Component
}

func (o *OpeningCharges) GetBalance() int {
	return o.component.GetBalance() - OPENING_CHARGE
}
