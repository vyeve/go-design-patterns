package factory

import "fmt"

// DebitCardPM is an example of PaymentMethod interface
type DebitCardPM struct{}

// Pay method satisfies to PaymentMethod interface
func (dc *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
