package domain

import "testing"

func TestNewDeposit_WhenWalletIDIsEmpty_ShouldReturnError(t *testing.T) {
	money, error := NewMoney(1000, "USD")
	if error != nil {
		t.Fatalf("unexpected error creating money: %v", error)
	}
	_, error = NewDeposit("deposit-1", "", money)

	if error == nil {
		t.Errorf("expected error when walletID is empty")
	}
}

func TestNewDeposit_WhenMoneyIsValid_ShouldCreateDeposit(t *testing.T) {
	money, error := NewMoney(1000, "USD")
	if error != nil {
		t.Fatalf("unexpected error creating money: %v", error)
	}

	deposit, err := NewDeposit("deposit-1", "wallet-1", money)
	if err != nil {
		t.Fatalf("unexpected error creating deposit: %v", err)
	}

	if deposit.WalletID != "wallet-1" {
		t.Errorf("expected walletID 'wallet-1', got '%s'", deposit.WalletID)
	}

	if deposit.Money.Amount != 1000 {
		t.Errorf("expected amount 1000, got %d", deposit.Money.Amount)
	}

	if deposit.Money.Currency != "USD" {
		t.Errorf("expected currency USD, got %s", deposit.Money.Currency)
	}
}
