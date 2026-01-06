package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type DepositFundsUseCase struct {
	txFactory TxFactory
	now       func() time.Time
}

func NewDepositFundsUseCase(txFactory TxFactory) *DepositFundsUseCase {
	return &DepositFundsUseCase{
		txFactory: txFactory,
		now:       func() time.Time { return time.Now().UTC() },
	}
}

type DepositFundsInput struct {
	WalletID string
	Amount   int64
	Currency string
}

type DepositFundsOutput struct {
	DepositID string
}

func (uc *DepositFundsUseCase) Execute(input DepositFundsInput) (DepositFundsOutput, error) {
	money, err := domain.NewMoney(input.Amount, input.Currency)
	if err != nil {
		return DepositFundsOutput{}, err
	}

	deposit, err := domain.NewDeposit(uuid.NewString(), input.WalletID, money)
	if err != nil {
		return DepositFundsOutput{}, err
	}

	tx, err := uc.txFactory.Begin()
	if err != nil {
		return DepositFundsOutput{}, err
	}
	defer tx.Rollback()

	depositRepo := uc.txFactory.NewDepositRepository(tx)
	ledgerRepo := uc.txFactory.NewLedgerRepository(tx)
	outboxRepo := uc.txFactory.NewOutboxRepository(tx)

	if err := depositRepo.Save(deposit); err != nil {
		return DepositFundsOutput{}, err
	}

	now := uc.now()

	credit, err := domain.NewLedgerEntry(
		deposit.WalletID,
		deposit.Money.Amount,
		deposit.Money.Currency,
		"credit",
		deposit.ID,
		now,
	)
	if err != nil {
		return DepositFundsOutput{}, err
	}

	debit, err := domain.NewLedgerEntry(
		"cash-in",
		deposit.Money.Amount,
		deposit.Money.Currency,
		"debit",
		deposit.ID,
		now,
	)
	if err != nil {
		return DepositFundsOutput{}, err
	}

	if err := ledgerRepo.Save(credit); err != nil {
		return DepositFundsOutput{}, err
	}
	if err := ledgerRepo.Save(debit); err != nil {
		return DepositFundsOutput{}, err
	}

	if err := outboxRepo.Save("deposit", deposit.ID, "deposit_processed", deposit); err != nil {
		return DepositFundsOutput{}, err
	}

	if err := tx.Commit(); err != nil {
		return DepositFundsOutput{}, err
	}

	return DepositFundsOutput{DepositID: deposit.ID}, nil
}
