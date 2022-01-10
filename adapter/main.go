package main

func main() {
	bank := NewBank()
	bank.ProcessPayment(105364, 1847630, 1_000_00)

	bankAdapter := NewBankAdapter(bank)
	bankAdapter.Send("alex@merren.com", "Frankie@dutoit.com", 1_000_00)

	paypal := NewPaypal()
	paypal.Send("alex@merren.com", "Frankie@dutoit.com", 1_000_00)
}
