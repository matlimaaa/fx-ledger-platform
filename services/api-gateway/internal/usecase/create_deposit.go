package usecase

import (
	"github.com/google/uuid"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type CreateDepositInput struct {
	WalletID string
	Amount   int64
	Currency string
}

type CreateDepositOutput struct {
	DepositID string
	Deposit   domain.Deposit
}

type DepositRepository interface {
	Save(deposit domain.Deposit) error
}

type CreateDepositUseCase struct {
	repo DepositRepository
}

func NewCreateDepositUseCase(repo DepositRepository) *CreateDepositUseCase {
	return &CreateDepositUseCase{repo: repo}
}

func (uc *CreateDepositUseCase) Execute(input CreateDepositInput) (CreateDepositOutput, error) {
	money, err := domain.NewMoney(input.Amount, input.Currency)
	if err != nil {
		return CreateDepositOutput{}, err
	}

	deposit, err := domain.NewDeposit(
		uuid.NewString(),
		input.WalletID,
		money,
	)

	if err != nil {
		return CreateDepositOutput{}, err
	}

	if err := uc.repo.Save(deposit); err != nil {
		return CreateDepositOutput{}, err
	}

	return CreateDepositOutput{
		DepositID: deposit.ID,
		Deposit:   deposit,
	}, nil
}
