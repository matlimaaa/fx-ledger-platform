package domain

import (
	"errors"
	"time"
)

type Deposit struct {
	ID string
	WalletID string
	Money Money
	CreatedAt time.Time
}

func NewDeposit(id, walletID string, money Money) (Deposit, error) {
	if walletID == "" {
		return Deposit{}, errors.New("wallet id is required")
	}

	return Deposit{
		ID: id,
		WalletID: walletID,
		Money: money,
		CreatedAt: time.Now(),
	}, nil
}