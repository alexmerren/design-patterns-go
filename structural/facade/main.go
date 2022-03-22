package main

func main() {
	atm := NewATMFacade("00001", 1932)
	atm.WithdrawCash("00001", 1932, 95_00)
}
