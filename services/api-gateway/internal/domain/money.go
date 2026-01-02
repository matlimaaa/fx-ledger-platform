package domain

import "errors"

type Money struct {
	Amount int64
	Currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount <= 0 {
		return Money{}, errors.New("amount must be greater than zero")
	}

	if currency == "" {
		return Money{}, errors.New("currency is required")
	}

	return Money{
		Amount: amount,
		Currency: currency,
	}, nil
}