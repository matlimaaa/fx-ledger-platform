package postgres

import (
	"database/sql"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/usecase"
)

type TxFactory struct {
	db *sql.DB
}

func NewTxFactory(db *sql.DB) *TxFactory {
	return &TxFactory{db: db}
}

func (f *TxFactory) Begin() (*sql.Tx, error) {
	return f.db.Begin()
}

func (f *TxFactory) NewDepositRepository(tx *sql.Tx) usecase.DepositRepositoryTx {
	return &DepositRepositoryTx{tx: tx}
}

func (f *TxFactory) NewLedgerRepository(tx *sql.Tx) usecase.LedgerRepositoryTx {
	return &LedgerRepository{tx: tx}
}

func (f *TxFactory) NewOutboxRepository(tx *sql.Tx) usecase.OutboxRepositoryTx {
	return &OutboxRepository{tx: tx}
}
