package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("a payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("the cash payment method wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("a payment method of type 'DebitCard' must exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("the debit card payment method wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	id := 20
	_, err := GetPaymentMethod(id)
	if err == nil {
		t.Errorf("a payment method with ID %d must return an error", id)
	}
	t.Logf("LOG: %v", err)

}
