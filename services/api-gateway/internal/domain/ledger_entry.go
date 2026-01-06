package domain

import (
	"errors"
	"time"
)

type LedgerEntry struct {
	WalletID    string
	Amount      int64
	Currency    string
	Direction   string
	ReferenceID string
	CreatedAt   time.Time
}

func NewLedgerEntry(
	walletID string,
	amount int64,
	currency string,
	direction string,
	referenceID string,
	createdAt time.Time,
) (LedgerEntry, error) {
	if walletID == "" {
		return LedgerEntry{}, errors.New("wallet id is required")
	}

	if amount <= 0 {
		return LedgerEntry{}, errors.New("amount must be greater than zero")
	}

	if currency == "" {
		return LedgerEntry{}, errors.New("currency is required")
	}

	if direction != "credit" && direction != "debit" {
		return LedgerEntry{}, errors.New("direction must be credit or debit")
	}

	if referenceID == "" {
		return LedgerEntry{}, errors.New("reference id is required")
	}

	return LedgerEntry{
		WalletID:    walletID,
		Amount:      amount,
		Currency:    currency,
		Direction:   direction,
		ReferenceID: referenceID,
		CreatedAt:   createdAt,
	}, nil
}
