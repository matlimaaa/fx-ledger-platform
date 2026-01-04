package usecase

import (
	"testing"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type FakeDepositRepository struct {
	saved []domain.Deposit
}

func (f *FakeDepositRepository) Save(d domain.Deposit) error {
	f.saved = append(f.saved, d)
	return nil
}

func TestCreateDepositUseCase(t *testing.T) {
	repo := &FakeDepositRepository{}
	uc := NewCreateDepositUseCase(repo)

	output, err := uc.Execute(CreateDepositInput{
		WalletID: "wallet-1",
		Amount:   1000,
		Currency: "USD",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if output.DepositID == "" {
		t.Errorf("expected deposit id to be generated")
	}

	if len(repo.saved) != 1 {
		t.Errorf("expected 1 deposit to be saved")
	}
}
