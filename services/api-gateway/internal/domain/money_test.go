package domain

import "testing"

func TestNewMoney(t *testing.T) {
	_, error := NewMoney(0, "USD")
	if error == nil {
		t.Errorf("expected error when amount is zero")
	}

	money, error := NewMoney(1000, "USD")
	if error != nil {
		t.Errorf("unexpected error: %v", error)
	}

	if money.Amount != 1000 {
		t.Errorf("expected amount 1000, got %d", money.Amount)
	}
}