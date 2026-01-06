package usecase

import (
	"database/sql"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type DepositRepositoryTx interface {
	Save(deposit domain.Deposit) error
}

type LedgerRepositoryTx interface {
	Save(entry domain.LedgerEntry) error
}

type OutboxRepositoryTx interface {
	Save(
		aggregate string,
		aggregateID string,
		eventType string,
		payload any,
	) error
}

type TxFactory interface {
	Begin() (*sql.Tx, error)

	NewDepositRepository(tx *sql.Tx) DepositRepositoryTx
	NewLedgerRepository(tx *sql.Tx) LedgerRepositoryTx
	NewOutboxRepository(tx *sql.Tx) OutboxRepositoryTx
}
