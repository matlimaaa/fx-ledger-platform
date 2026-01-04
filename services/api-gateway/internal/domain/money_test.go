package domain

import "testing"

func TestNewMoney(t *testing.T) {
	_, err := NewMoney(0, "USD")
	if err == nil {
		t.Errorf("expected error when amount is zero")
	}

	money, err := NewMoney(1000, "USD")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if money.Amount != 1000 {
		t.Errorf("expected amount 1000, got %d", money.Amount)
	}
}
