package main

import "fmt"

func main() {
	account := NewBankAccount("test", 129_00)
	final := &InterestCharges{
		component: &OpeningCharges{
			component: &OverdraftCharges{
				component: account,
			},
		},
	}

	fmt.Println(final.GetBalance())
}
