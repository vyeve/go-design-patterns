/*
The Factory method pattern abstracts the user from knowledge of the struct he needs
to achieve for a specific purpose,such as retrieving some value, maybe from a web, service
or a database. The user only needs an interface that provides him this value. By delegating
this decision to a Factory, this Factory can provide an interface that fits the user needs.
*/

package factory

import (
	"fmt"
)

// PaymentMethod interface represents method Pay
type PaymentMethod interface {
	Pay(amount float32) string
}

// different types of payments
const (
	Cash = iota
	DebitCard
)

// GetPaymentMethod return an object that satisfies PaymentMethod interface
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method with ID %d not recognized.\n", m)
	}
}
