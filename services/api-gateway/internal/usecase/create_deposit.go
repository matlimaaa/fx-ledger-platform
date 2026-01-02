package usecase

import "github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"

type DepositRepository interface {
	Save(deposit domain.Deposit) error
}

type CreateDepositUseCase struct {
	repo DepositRepository
}

func NewCreateDepositUseCase(repo DepositRepository) *CreateDepositUseCase {
	return &CreateDepositUseCase{repo: repo}
}

func (uc *CreateDepositUseCase) Execute(deposit domain.Deposit) error {
	return uc.repo.Save(deposit)
}
