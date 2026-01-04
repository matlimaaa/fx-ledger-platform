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
	money, error := domain.NewMoney(input.Amount, input.Currency)
	if error != nil {
		return CreateDepositOutput{}, error
	}

	deposit, error := domain.NewDeposit(
		uuid.NewString(),
		input.WalletID,
		money,
	)

	if error != nil {
		return CreateDepositOutput{}, error
	}

	if error := uc.repo.Save(deposit); error != nil {
		return CreateDepositOutput{}, error
	}

	return CreateDepositOutput{
		DepositID: deposit.ID,
	}, nil
}
