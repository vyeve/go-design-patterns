package creational

import "fmt"

// CashPM is an example of PaymentMethod interface
type CashPM struct{}

// Pay method satisfies to PaymentMethod interface
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}
