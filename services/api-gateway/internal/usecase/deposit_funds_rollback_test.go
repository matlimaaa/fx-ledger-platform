package usecase_test

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/adapters/postgres"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/usecase"
)

type failingOutboxRepository struct{}

func (f *failingOutboxRepository) Save(aggregate, aggregateID, eventType string, payload any) error {
	return errors.New("outbox failure")
}

type failingTxFactory struct {
	real usecase.TxFactory
}

func (f *failingTxFactory) Begin() (*sql.Tx, error) {
	return f.real.Begin()
}

func (f *failingTxFactory) NewDepositRepository(tx *sql.Tx) usecase.DepositRepositoryTx {
	return f.real.NewDepositRepository(tx)
}

func (f *failingTxFactory) NewLedgerRepository(tx *sql.Tx) usecase.LedgerRepositoryTx {
	return f.real.NewLedgerRepository(tx)
}

func (f *failingTxFactory) NewOutboxRepository(tx *sql.Tx) usecase.OutboxRepositoryTx {
	return &failingOutboxRepository{}
}

func TestDepositFunds_RollbackOnOutboxFailure(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://fx:fx@localhost:5432/fx_ledger?sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	_, _ = db.Exec("DELETE FROM outbox_events")
	_, _ = db.Exec("DELETE FROM ledger_entries")
	_, _ = db.Exec("DELETE FROM deposits")

	realFactory := postgres.NewTxFactory(db)

	var _ usecase.TxFactory = realFactory

	factory := &failingTxFactory{real: realFactory}
	uc := usecase.NewDepositFundsUseCase(factory)

	_, err = uc.Execute(usecase.DepositFundsInput{
		WalletID: "wallet-rollback",
		Amount:   1000,
		Currency: "USD",
	})
	require.Error(t, err)

	assertCount(t, db, "deposits", 0)
	assertCount(t, db, "ledger_entries", 0)
	assertCount(t, db, "outbox_events", 0)
}

func assertCount(t *testing.T, db *sql.DB, table string, want int) {
	t.Helper()
	var got int
	err := db.QueryRow("SELECT count(*) FROM " + table).Scan(&got)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
